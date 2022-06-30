/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

import type {Request} from 'express';
import type {UserModel} from '../../shared/sequelize_models/models/user';
import type {UserRawType} from '../../shared/sequelize_models/models/user';

import bcrypt from 'bcryptjs';
import querystring from 'querystring';
import {AccessRoles} from '../../shared/roles';
import {User} from '../../shared/sequelize_models';
import {format, parse} from 'url';
import {injectOrganizationParams} from './organization';
import {validate as validateEmail} from 'email-validator';

const SALT_GEN_ROUNDS = 10;
const MIN_PASSWORD_LENGTH = 10;

const FIELD_MAP = {
  email: 'email',
  networkIDs: 'networkIDs',
  organization: 'organization',
  password: 'password',
  role: 'role',
};

export function addQueryParamsToUrl(
  url: string,
  params: {[key: string]: any},
): string {
  const parsedUrl = parse(url, true /* parseQueryString */);
  if (params) {
    parsedUrl.search = querystring.stringify({
      ...parsedUrl.query,
      ...params,
    });
  }
  return format(parsedUrl);
}

export async function getUserFromRequest(
  req: Request,
  email: string,
): Promise<UserModel | null | undefined> {
  const where = await injectOrganizationParams(req, {email});
  return await User.findOne({where});
}

export async function getPropsToUpdate(
  allowedProps: ReadonlyArray<keyof typeof FIELD_MAP>,
  body: Partial<UserRawType>,
  organizationInjector: (params: {
    [key: string]: any;
  }) => Promise<{[key: string]: any; organization?: string}>,
): Promise<Partial<UserRawType>> {
  allowedProps = allowedProps.filter(prop =>
    User.rawAttributes.hasOwnProperty(FIELD_MAP[prop]),
  );
  const userProperties: Record<string, number | string | Array<string>> = {};
  for (const prop of allowedProps) {
    if (body.hasOwnProperty(prop)) {
      switch (prop) {
        case 'email':
          const emailUnsafe = body[prop];
          if (typeof emailUnsafe !== 'string' || !validateEmail(emailUnsafe)) {
            throw new Error('Please enter a valid email');
          }
          const email = emailUnsafe.toLowerCase();

          // Check if user exists
          const where = await organizationInjector({email});
          if (await User.findOne({where})) {
            throw new Error(`${email} already exists`);
          }
          userProperties[prop] = email;
          break;
        case 'password':
          userProperties[prop] = await validateAndHashPassword(
            String(body[prop]),
          );
          break;
        case 'role':
          userProperties[prop] =
            body[prop] === AccessRoles.SUPERUSER
              ? AccessRoles.SUPERUSER
              : body[prop] === AccessRoles.READ_ONLY_USER
              ? AccessRoles.READ_ONLY_USER
              : AccessRoles.USER;
          break;
        case 'networkIDs':
          const networkIDsunsafe = body[prop];
          if (Array.isArray(networkIDsunsafe)) {
            const networkIDs: Array<string> = networkIDsunsafe.map(it => {
              if (typeof it !== 'string') {
                throw new Error('Please enter valid network IDs');
              }
              return it;
            });
            userProperties[prop] = networkIDs;
            break;
          }
          throw new Error('Please enter valid network IDs');
        case 'organization':
          if (typeof body[prop] !== 'string') {
            throw new Error('Invalid Organization!');
          }
          userProperties[prop] = body[prop] as string;
          break;
        default:
          userProperties[prop] = body[prop];
          break;
      }
    }
  }
  return userProperties;
}

export async function validateAndHashPassword(password: string) {
  if (
    typeof password !== 'string' ||
    password === '' ||
    password.length < MIN_PASSWORD_LENGTH
  ) {
    throw new Error(
      `Password must contain at least ${MIN_PASSWORD_LENGTH} characters`,
    );
  }

  const salt = await bcrypt.genSalt(SALT_GEN_ROUNDS);
  return await bcrypt.hash(password, salt);
}

/**
 * Copyright 2022 The Magma Authors.
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
 * @flow strict-local
 * @format
 */

// $FlowFixMe[cannot-resolve-module] for TypeScript migration
import useAxios from '../../hooks/useAxios';

// $FlowIgnore because responses contains multiple types
export function mockUseAxios(responses: {[url: string]: {data: any}}) {
  // $FlowIgnore
  useAxios.mockImplementation(config => {
    const response = responses[config.url];
    if ('onResponse' in config) {
      if (response) {
        config.onResponse(response);
        delete responses[config.url];
      }
      return {};
    } else {
      return {response};
    }
  });
}

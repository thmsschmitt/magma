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
 */

import * as React from 'react';
import ConfigEditor from './ConfigEditor';
import Grid from '@mui/material/Grid';
import TextField from '@mui/material/TextField';
import type {EditorProps} from './ConfigEditor';
import type {ReceiverPagerDutyConfig} from '../../AlarmAPIType';

export default function PagerDutyConfigEditor({
  config,
  onUpdate,
  ...props
}: EditorProps<ReceiverPagerDutyConfig>) {
  return (
    <ConfigEditor
      {...props}
      RequiredFields={
        <>
          <Grid item>
            <TextField
              variant="standard"
              required
              label="Description"
              value={config.description}
              onChange={e => onUpdate({description: e.target.value})}
              fullWidth
            />
          </Grid>
          <Grid item>
            <TextField
              variant="standard"
              required
              label="Severity"
              value={config.severity}
              onChange={e => onUpdate({severity: e.target.value})}
              fullWidth
            />
          </Grid>
          <Grid item>
            <TextField
              variant="standard"
              required
              label="Url"
              placeholder="Ex: webhook.example.com"
              value={config.url}
              onChange={e => onUpdate({url: e.target.value})}
              fullWidth
            />
          </Grid>
          <Grid item>
            <TextField
              variant="standard"
              required
              label="Routing_key"
              value={config.routing_key}
              onChange={e => onUpdate({routing_key: e.target.value})}
              fullWidth
            />
          </Grid>
          <Grid item>
            <TextField
              variant="standard"
              required
              label="Service_key"
              value={config.service_key}
              onChange={e => onUpdate({service_key: e.target.value})}
              fullWidth
            />
          </Grid>
          <Grid item>
            <TextField
              variant="standard"
              required
              label="Client"
              value={config.client}
              onChange={e => onUpdate({client: e.target.value})}
              fullWidth
            />
          </Grid>
          <Grid item>
            <TextField
              variant="standard"
              required
              label="Client Url"
              value={config.client_url}
              onChange={e => onUpdate({client_url: e.target.value})}
              fullWidth
            />
          </Grid>
        </>
      }
      data-testid="pager-duty-config-editor"
    />
  );
}

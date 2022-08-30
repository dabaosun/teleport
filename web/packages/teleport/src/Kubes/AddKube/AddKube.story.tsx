/*
Copyright 2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import React from 'react';
import { Attempt } from 'shared/hooks/useAttemptNext';

import { AddKube, Props } from './AddKube';
import { State } from './useAddKube';

export default {
  title: 'Teleport/Kubes/Add',
};

export const Loaded = () => <AddKube {...props} />;

export const TokenGenerated = () => (
  <AddKube
    {...props}
    token={{ id: 'some token', expiry: null, expiryText: '4 hours' }}
  />
);

export const Processing = () => (
  <AddKube {...props} attempt={{ status: 'processing' }} />
);

export const Failed = () => (
  <AddKube
    {...props}
    attempt={{ status: 'failed', statusText: 'some error message' }}
  />
);

const props: Props & State = {
  onClose: () => null,
  createToken: () => Promise.resolve(null),
  attempt: {
    status: 'success',
    statusText: '',
  },
  token: null,
  version: '10.0.0',
};

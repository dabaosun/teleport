/**
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import {
  DefaultTab,
  LabelsViewMode,
  ViewMode,
} from 'gen-proto-ts/teleport/userpreferences/v1/unified_resource_preferences_pb';

import { AssistViewMode } from 'gen-proto-ts/teleport/userpreferences/v1/assist_pb';

import { UserPreferences } from 'gen-proto-ts/teleport/userpreferences/v1/userpreferences_pb';

import { ClusterUserPreferences } from 'gen-proto-ts/teleport/userpreferences/v1/cluster_preferences_pb';

import { Theme } from 'gen-proto-ts/teleport/userpreferences/v1/theme_pb';

import cfg from 'teleport/config';
import api from 'teleport/services/api';

import { KeysEnum } from '../storageService';

export async function getUserPreferences() {
  const res: UserPreferences = await api.get(cfg.api.userPreferencesPath);

  return res;
}

export async function getUserClusterPreferences(clusterId: string) {
  return await api
    .get(cfg.getUserClusterPreferencesUrl(clusterId))
    .then(res => {
      // TODO (avatus) DELETE IN 16
      // this item is used to disabled the pinned resources button if they
      // haven't upgraded to 14.1.0 yet. Anything lower than 14 doesn't matter
      // because the unified resource view isn't enabled so pinning isn't there either
      localStorage.removeItem(KeysEnum.PINNED_RESOURCES_NOT_SUPPORTED);
      return res;
    })
    .catch(res => {
      if (res.response?.status === 403 || res.response?.status === 404) {
        localStorage.setItem(KeysEnum.PINNED_RESOURCES_NOT_SUPPORTED, 'true');
        // we handle this null error in the user context where we cache cluster
        // preferences. We want to fail gracefully here and use our "not supported"
        // message instead.
        return null;
      }
      // return all other errors here
      return res;
    });
}

export function updateUserClusterPreferences(
  clusterId: string,
  preferences: UserPreferences
) {
  return api.put(cfg.getUserClusterPreferencesUrl(clusterId), preferences);
}

export function updateUserPreferences(preferences: Partial<UserPreferences>) {
  return api.put(cfg.api.userPreferencesPath, preferences);
}

export function makeDefaultUserPreferences(): UserPreferences {
  return {
    theme: Theme.LIGHT,
    assist: {
      viewMode: AssistViewMode.DOCKED,
      preferredLogins: [],
    },
    onboard: {
      preferredResources: [],
      marketingParams: {
        campaign: '',
        source: '',
        medium: '',
        intent: '',
      },
    },
    unifiedResourcePreferences: {
      defaultTab: DefaultTab.ALL,
      viewMode: ViewMode.CARD,
      labelsViewMode: LabelsViewMode.COLLAPSED,
    },
    clusterPreferences: makeDefaultUserClusterPreferences(),
  };
}

export function makeDefaultUserClusterPreferences(): ClusterUserPreferences {
  return {
    pinnedResources: {
      resourceIds: [],
    },
  };
}

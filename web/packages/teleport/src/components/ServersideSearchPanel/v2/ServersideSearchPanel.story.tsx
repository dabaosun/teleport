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

import React, { useState } from 'react';

import { SearchInput } from './SearchInput';
import { ServersideSearchPanel } from './ServersideSearchPanel';

export default { title: 'Shared/Search/V2' };

export const Input = () => {
  const [search, setSearch] = useState<string>();

  return (
    <SearchInput
      dataType={'sunclasses'}
      searchValue={search}
      setSearchValue={setSearch}
    />
  );
};

export const ServersidePanel = () => {
  const [search, setSearch] = useState<string>();
  const [advanced, setAdvanced] = useState<boolean>();

  return (
    <ServersideSearchPanel
      dataType={'phone'}
      searchString={search}
      setSearchString={setSearch}
      isAdvancedSearch={advanced}
      setIsAdvancedSearch={setAdvanced}
      onSubmitSearch={() => {}}
    />
  );
};

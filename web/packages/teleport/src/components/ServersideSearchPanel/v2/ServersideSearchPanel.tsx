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

import React from 'react';
import { Flex } from 'design';

import { AdvancedSearchToggle } from 'shared/components/AdvancedSearchToggle';

import useServersideSearchPanel, {
  HookProps,
  SearchPanelState,
} from 'teleport/components/ServersideSearchPanel/useServerSideSearchPanel';

import { SearchInput } from './SearchInput';

export default function Container(props: HookProps & { dataType: string }) {
  const state = useServersideSearchPanel(props);
  return <ServersideSearchPanel dataType={props.dataType} {...state} />;
}

export function ServersideSearchPanel({
  dataType,
  isAdvancedSearch,
  onSubmitSearch,
  searchString,
  setIsAdvancedSearch,
  setSearchString,
}: SearchPanelState & { dataType: string }) {
  function onToggle() {
    setIsAdvancedSearch(wasAdvancedSearch => !wasAdvancedSearch);
  }

  return (
    <Flex
      as="form"
      data-testid="form"
      className="SearchPanel"
      onSubmit={onSubmitSearch}
      mb={3}
    >
      <SearchInput
        dataType={dataType}
        searchValue={searchString}
        setSearchValue={setSearchString}
      >
        <AdvancedSearchToggle
          isToggled={isAdvancedSearch}
          onToggle={onToggle}
          px={4}
        />
      </SearchInput>
    </Flex>
  );
}

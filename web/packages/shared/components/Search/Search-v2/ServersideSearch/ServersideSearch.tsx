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

import { Box, Flex } from 'design';
import { StyledPanel } from 'design/DataTable';
import InputSearch from 'design/DataTable/InputSearch';

import { PageIndicatorText } from 'shared/components/Search';

import { AdvancedSearchToggle } from '../AdvancedSearchToggle';
import { ServersideSearchProps, ServersideSearchState } from '../types';
import useServersideSearch from '../hooks/useServerSideSearch';

export default function Container(props: ServersideSearchProps) {
  const { pageIndicators, disabled, ...hookProps } = props;
  const state = useServersideSearch(hookProps);
  return (
    <ServersideSearch
      {...state}
      pageIndicators={pageIndicators}
      disabled={disabled}
    />
  );
}

export function ServersideSearch({
  searchString,
  setSearchString,
  isAdvancedSearch,
  setIsAdvancedSearch,
  onSubmitSearch,
  pageIndicators,
  disabled = false,
}: ServersideSearchState) {
  function onToggle() {
    setIsAdvancedSearch(!isAdvancedSearch);
  }

  return (
    <StyledPanel
      data-testid="form"
      as="form"
      onSubmit={onSubmitSearch}
      borderTopLeftRadius={3}
      borderTopRightRadius={3}
      style={disabled ? { pointerEvents: 'none', opacity: '0.5' } : {}}
    >
      <Flex justifyContent="space-between" alignItems="center" width="100%">
        <Flex style={{ width: '70%' }} alignItems="center">
          <Box width="100%" mr={3}>
            <InputSearch
              ariaLabel="search input"
              searchValue={searchString}
              setSearchValue={setSearchString}
            >
              <AdvancedSearchToggle
                isToggled={isAdvancedSearch}
                onToggle={onToggle}
                px={4}
              />
            </InputSearch>
          </Box>
        </Flex>
        <Flex>
          <PageIndicatorText
            from={pageIndicators.from}
            to={pageIndicators.to}
            count={pageIndicators.totalCount}
          />
        </Flex>
      </Flex>
    </StyledPanel>
  );
}

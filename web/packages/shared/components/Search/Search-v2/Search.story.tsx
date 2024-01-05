import React, { useState } from 'react';

import { AdvancedSearchInput } from './AdvancedSearchInput';
import { SearchInput } from './SearchInput';
import { AdvancedSearchToggle } from './AdvancedSearchToggle';
import { ServersideSearch } from './ServersideSearch/ServersideSearch';

export default { title: 'Shared/Search/V2' };

export const Input = () => {
  const [search, setSearch] = useState<string>();

  return (
    <SearchInput
      dataType={'sUnGLASSes'}
      searchValue={search}
      setSearchValue={setSearch}
    />
  );
};

export const AdvancedInput = () => {
  const [search, setSearch] = useState<string>();
  const [advanced, setAdvanced] = useState<boolean>();

  return (
    <AdvancedSearchInput
      dataType={'KeYs'}
      searchString={search}
      setSearchString={setSearch}
      isAdvancedSearch={advanced}
      setIsAdvancedSearch={setAdvanced}
      onSubmitSearch={() => {}}
    />
  );
};

export const AdvancedToggle = () => {
  const [toggled, setToggled] = useState(false);

  function toggle(): void {
    setToggled(wasToggled => !wasToggled);
  }

  return (
    <AdvancedSearchToggle
      isToggled={toggled}
      onToggle={toggle}
    />
  );
};

export const Serverside = () => {
  const [search, setSearch] = useState<string>();
  const [advanced, setAdvanced] = useState<boolean>();

  return (
    <ServersideSearch
      dataType={''}
      searchString={search}
      setSearchString={setSearch}
      isAdvancedSearch={advanced}
      setIsAdvancedSearch={setAdvanced}
      onSubmitSearch={() => {}}
      pageIndicators={{ to: 0, from: 0, totalCount: 0 }}
    />
  );
};

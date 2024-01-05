import React from 'react';

import { SearchPagination } from './SearchPagination';
import { SearchPanel } from './SearchPanel';

export default { title: 'Shared/Search/V1' };

export const Pagination = () => {
  return <SearchPagination prevPage={() => {}} nextPage={() => {}} />;
};

export const Panel = () => {
  return (
    <SearchPanel
      updateQuery={() => {}}
      updateSearch={() => {}}
      pageIndicators={{ from: 0, to: 0, total: 0 }}
      filter={{
        query: '',
        search: '',
        sort: { fieldName: '', dir: 'ASC' },
        limit: 0,
        startKey: '',
        pinnedOnly: false,
        searchAsRoles: '',
        kinds: [],
      }}
      showSearchBar={false}
      disableSearch={false}
    />
  );
};

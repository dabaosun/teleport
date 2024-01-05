import { Dispatch, FormEvent, JSX, SetStateAction } from 'react';

export type AdvancedSearchToggleProps = {
  isToggled: boolean;
  onToggle(): void;
  px?: number | string;
  gap?: number;
  className?: string;
};

export type SearchInputProps = {
  dataType: string;
  searchValue: string;
  setSearchValue: Dispatch<SetStateAction<string>>;
  children?: JSX.Element;
  ariaLabel?: string;
};

export type ServerSidePassThroughProps = {
  pageIndicators: PageIndicators;
  disabled?: boolean;
};

export type UseServersideSearchProps = {
  dataType: string;
  pathname: string;
  replaceHistory: (path: string) => void;
  params: SearchFilter;
  setParams: (params: SearchFilter) => void;
};

export type ServersideSearchProps = UseServersideSearchProps &
  ServerSidePassThroughProps;

export type SearchState = {
  dataType: string;
  searchString: string;
  setSearchString: Dispatch<SetStateAction<string>>;
  isAdvancedSearch: boolean;
  setIsAdvancedSearch: Dispatch<SetStateAction<boolean>>;
  onSubmitSearch: (e: FormEvent<HTMLFormElement>) => void;
};

export type ServersideSearchState = SearchState & ServerSidePassThroughProps;

export type SearchFilter = {
  /** query is query expression using the predicate language. */
  query?: string;
  /** search contains search words/phrases separated by space. */
  search?: string;
  sort?: SortType;
  limit?: number;
  startKey?: string;
  pinnedOnly?: boolean;
  searchAsRoles?: '' | 'yes';
  // TODO(bl-nero): Remove this once filters are expressed as advanced search.
  kinds?: string[];
};

export type SortType = {
  fieldName: string;
  dir: SortDir;
};

export type SortDir = 'ASC' | 'DESC';

/** Contains the values needed to display 'Showing X - X of X' on the top right of the table. */
export type PageIndicators = {
  /** The position of the first item on the page relative to all items. */
  from: number;
  /** The position of the last item on the page relative to all items. */
  to: number;
  /** The total number of all items. */
  totalCount: number;
};

import { render, screen } from 'design/utils/testing';

import { AdvancedSearchInput } from './AdvancedSearchInput';
import { SearchState } from './types';

test('renders search input with advanced toggle', () => {
  const props: SearchState = {
    dataType: '',
    searchString: '',
    setSearchString: () => {},
    isAdvancedSearch: false,
    setIsAdvancedSearch: () => {},
    onSubmitSearch: () => {},
  };
  render(<AdvancedSearchInput {...props} />);

  expect(screen.getByLabelText('search input')).toBeInTheDocument();
  expect(screen.getByText('Advanced')).toBeInTheDocument();
});

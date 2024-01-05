import { fireEvent, render, screen, userEvent } from 'design/utils/testing';

import { ServersideSearchState } from '../types';

import { ServersideSearch } from './ServersideSearch';

test('renders input with advanced toggle, calls cbs', async () => {
  const mockSetAdvanced = jest.fn();
  const mockSetSearch = jest.fn();
  const mockOnSubmit = jest.fn();

  const props: ServersideSearchState = {
    dataType: '',
    searchString: '',
    isAdvancedSearch: true,
    setSearchString: mockSetSearch,
    setIsAdvancedSearch: mockSetAdvanced,
    onSubmitSearch: mockOnSubmit,
    pageIndicators: {
      from: 0,
      to: 0,
      totalCount: 0,
    },
  };
  render(<ServersideSearch {...props} />);

  expect(screen.getByLabelText('search input')).toBeInTheDocument();
  expect(screen.getByRole('checkbox')).toBeInTheDocument();

  await userEvent.click(screen.getByRole('checkbox'));
  expect(mockSetAdvanced).toHaveBeenCalled();

  await userEvent.type(screen.getByLabelText('search input'), 'a{enter}');
  expect(mockSetSearch).toHaveBeenCalledWith('a');

  fireEvent.submit(screen.getByTestId('form'));
  expect(mockOnSubmit).toHaveBeenCalled();
});

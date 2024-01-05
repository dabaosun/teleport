import React from 'react';

import { render, screen, userEvent } from 'design/utils/testing';

import { SearchInput } from './SearchInput';
import { SearchInputProps } from './types';

test('renders resource type & calls cb', async () => {
  const mockSet = jest.fn();
  const props: SearchInputProps = {
    dataType: 'bOTS',
    searchValue: '',
    setSearchValue: mockSet,
    children: <>Hello</>,
    ariaLabel: 'search input',
  };

  render(<SearchInput {...props} />);
  expect(screen.getByLabelText('search input')).toHaveAttribute(
    'placeholder',
    'Search for Bots...'
  );

  await userEvent.type(screen.getByLabelText('search input'), 'admin');
  expect(mockSet).toHaveBeenCalledTimes(5);
  expect(mockSet).toHaveBeenCalledWith('a');
  expect(mockSet).toHaveBeenCalledWith('d');
  expect(mockSet).toHaveBeenCalledWith('m');
  expect(mockSet).toHaveBeenCalledWith('i');
  expect(mockSet).toHaveBeenCalledWith('n');

  expect(screen.getByText('Hello')).toBeInTheDocument();
});

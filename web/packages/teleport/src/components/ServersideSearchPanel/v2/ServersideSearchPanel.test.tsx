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
import { fireEvent, render, screen, userEvent } from 'design/utils/testing';

import { SearchPanelState } from '../useServerSideSearchPanel';

import { ServersideSearchPanel } from './ServersideSearchPanel';

test('renders input with advanced toggle, calls cbs', async () => {
  const mockSetAdvanced = jest.fn();
  const mockSetSearch = jest.fn();
  const mockOnSubmit = jest.fn();

  const props: SearchPanelState & { dataType: string } = {
    dataType: '',
    searchString: '',
    isAdvancedSearch: true,
    setSearchString: mockSetSearch,
    setIsAdvancedSearch: mockSetAdvanced,
    onSubmitSearch: mockOnSubmit,
  };
  render(<ServersideSearchPanel {...props} />);

  expect(
    screen.getByRole('textbox', { name: 'search-input' })
  ).toBeInTheDocument();
  expect(screen.getByRole('checkbox')).toBeInTheDocument();

  await userEvent.click(screen.getByRole('checkbox'));
  expect(mockSetAdvanced).toHaveBeenCalled();

  await userEvent.type(screen.getByRole('textbox', { name: 'search-input' }), 'a{enter}');
  expect(mockSetSearch).toHaveBeenCalledWith('a');

  fireEvent.submit(screen.getByTestId('form'));
  expect(mockOnSubmit).toHaveBeenCalled();
});

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
    };

    render(<SearchInput {...props} />);
    expect(screen.getByRole('textbox', {name: 'search-input'})).toHaveAttribute(
        'placeholder',
        'Search for Bots...'
    );

    await userEvent.type(screen.getByRole('textbox', {name: 'search-input'}), 'admin');
    expect(mockSet).toHaveBeenCalledTimes(5);
    expect(mockSet).toHaveBeenCalledWith('a');
    expect(mockSet).toHaveBeenCalledWith('d');
    expect(mockSet).toHaveBeenCalledWith('m');
    expect(mockSet).toHaveBeenCalledWith('i');
    expect(mockSet).toHaveBeenCalledWith('n');

    expect(screen.getByText('Hello')).toBeInTheDocument();
});

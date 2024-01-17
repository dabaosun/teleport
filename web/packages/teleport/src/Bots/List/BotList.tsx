/**
 * Teleport
 * Copyright (C) 2024 Gravitational, Inc.
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

import Table, { LabelCell } from 'design/DataTable';

import React from 'react';

import { BotOptionsCell } from 'teleport/Bots/List/ActionCell';

import { BotListParams } from 'teleport/Bots/types';

export function BotList({ bots }: BotListParams) {
  return (
    <Table
      data={bots}
      columns={[
        {
          key: 'name',
          headerText: 'Bot Name',
          isSortable: true,
        },
        {
          key: 'kind',
          headerText: 'Type',
          isSortable: true,
        },
        {
          key: 'roles',
          headerText: 'Roles',
          isSortable: true,
          onSort: (a: string[], b: string[]) => {
            const aStr = a.toString();
            const bStr = b.toString();

            if (aStr < bStr) {
              return -1;
            }
            if (aStr > bStr) {
              return 1;
            }

            return 0;
          },
          render: ({ roles }) => <LabelCell data={roles} />,
        },
        {
          altKey: 'options-btn',
          render: bot => (
            <BotOptionsCell bot={bot} onEdit={() => {}} onDelete={() => {}} />
          ),
        },
      ]}
      emptyText="No Bots Found"
      isSearchable
      pagination={{ pageSize: 20 }}
    />
  );
}

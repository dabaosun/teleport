import { renderHook } from '@testing-library/react';

import useServersideSearch from 'shared/components/Search/Search-v2/hooks/useServerSideSearch';

test('initial state passes through', async () => {
  const { result } = renderHook(() =>
    useServersideSearch({
      dataType: 'wallet',
      pathname: '',
      replaceHistory: () => {},
      params: {},
      setParams: () => {},
    })
  );

  expect(result.current.dataType).toBe('wallet');
  expect(result.current.searchString).toBe('');
  expect(result.current.isAdvancedSearch).toBe(false);
});

test('parses param query term', async () => {
  const { result } = renderHook(() =>
    useServersideSearch({
      dataType: '',
      pathname: '',
      replaceHistory: () => {},
      params: {
        query: 'credit card',
      },
      setParams: () => {},
    })
  );

  expect(result.current.searchString).toBe('credit card');
});

test('parses param search term', async () => {
  const { result } = renderHook(() =>
    useServersideSearch({
      dataType: '',
      pathname: '',
      replaceHistory: () => {},
      params: {
        search: 'license',
      },
      setParams: () => {},
    })
  );

  expect(result.current.searchString).toBe('license');
});

import { render, screen, userEvent } from 'design/utils/testing';

import { AdvancedSearchToggleProps } from '../types';

import { AdvancedSearchToggle } from './AdvancedSearchToggle';

test('sets optional fields and calls cb', async () => {
  const mockToggle = jest.fn();
  const props: AdvancedSearchToggleProps = {
    isToggled: false,
    onToggle: mockToggle,
    px: 1, // 4px
    gap: 6, // 40px
    className: 'example-class',
  };

  const { container } = render(<AdvancedSearchToggle {...props} />);
  expect(screen.getByRole('checkbox')).not.toBeChecked();
  expect(container.firstChild).toHaveClass('example-class');
  expect(container.firstChild).toHaveStyle('gap: 40px');
  expect(container.firstChild).toHaveStyle('padding-left: 4px');
  expect(container.firstChild).toHaveStyle('padding-right: 4px');

  await userEvent.click(screen.getByRole('checkbox'));
  expect(mockToggle).toHaveBeenCalled();
});

test('checks toggle if true', async () => {
  const props: AdvancedSearchToggleProps = {
    isToggled: true,
    onToggle: () => {},
  };

  const { container } = render(<AdvancedSearchToggle {...props} />);
  expect(screen.getByRole('checkbox')).toBeChecked();
  expect(container.firstChild).toHaveStyle('gap: 8px'); // default
});

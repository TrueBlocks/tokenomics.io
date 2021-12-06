import { useEffect, useState, useCallback, useMemo } from 'react';

const navigationKeys = [
   'ArrowUp',
   'ArrowDown',
   'ArrowLeft',
   'PageUp',
   'ArrowRight',
   'PageDown',
   'End',
   'Home',
];

// Adds keyboard navigation logic. The hook returns row and page
// values that get increased/decreased when the arrow keys, page up,
// page down, home and end keys are pressed.
export function useKeyNav({ pageSize, maxItems }) {
  // Position in the dataset. In a set of maxItems items, it can be
  // an integer between 0 and maxItems.
  const [position, setPosition] = useState(0);
  // The navigation can be "turned off", which will not remove the
  // listeners and the logic will still work, but the state can be
  // presented as, for example, lack of highlight color.
  const [on, setOn] = useState(false);
  const page = useMemo(() => {
    return Math.floor(position / pageSize) + 1;
  }, [pageSize, position]);
  const row = useMemo(() => {
    return position % pageSize;
  }, [pageSize, position]);

  // Adds `addend` to `position`. To decrease the position value just use
  // a negative `addend`.
  const incrementPosition = useCallback((addend) => (currentPosition) => {
    return Math.max(0, Math.min(maxItems, currentPosition + addend));
  }, [maxItems]);
  // Handles the pressed key
  const listener = useCallback(({ code }) => {
    // Most of the time, arrow up or down will move by one item...
    let arrowUpAndDownAddend = 1;

    if (navigationKeys.includes(code) && on === false) {
      // ... but, if the navigation was "turned off" and the user
      // is turning it back on just now, we want to focus the first item
      // (which always has 0 index)
      arrowUpAndDownAddend = 0;
      setOn(true);
    }

    /* eslint-disable default-case */
    switch (code) {
      case 'ArrowUp':
        setPosition(incrementPosition(-arrowUpAndDownAddend));
        break;
      case 'ArrowDown':
        setPosition(incrementPosition(+arrowUpAndDownAddend));
        break;
      case 'ArrowLeft':
      case 'PageUp':
        setPosition(incrementPosition(-pageSize));
        break;
      case 'ArrowRight':
      case 'PageDown':
        setPosition(incrementPosition(+pageSize));
        break;
      case 'End':
        setPosition(maxItems);
        break;
      case 'Home':
        setPosition(0);
        break;
      case 'Escape':
        setPosition(0);
        setOn(false);
        break;
    }
  }, [on, incrementPosition, pageSize, maxItems]);

  useEffect(() => {
    window.addEventListener('keydown', listener);
    return () => window.removeEventListener('keydown', listener);
  }, [listener]);

  return {
    page,
    row,
    keyNavOn: on,
  };
}
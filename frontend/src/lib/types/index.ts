export type BoardStyle = 'numbers' | 'colors' | 'both';

export type Coordinate = {
  col: number;
  row: number;
};

export type DropdownOption = {
  label: string;
  onclick: () => void;
};

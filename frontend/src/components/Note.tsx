interface NoteProps {
  date: Date;
  description: string;
  value: number;
}

export function Note({ date, description, value }: NoteProps) {
  return <article />;
}

interface NoteProps {
  date: Date;
  description: string;
  value: number;
}

export function Note({ date, description, value }: NoteProps) {
  return (
    <article>
      <div>{date.toLocaleDateString()}</div>
      <div>{description}</div>
      <div>{value}</div>
    </article>
  );
}

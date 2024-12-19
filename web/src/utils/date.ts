export const Months = [
  "January",
  "February",
  "March",
  "April",
  "May",
  "June",
  "July",
  "August",
  "September",
  "October",
  "November",
  "December",
];

interface iOptionDateFormat {
  /**
   * @default " "
   */
  sparator?: string;

  /**
   * @default true
   */
  with_month?: boolean;
}

export function DateFormat(
  date: string | Date,
  option?: iOptionDateFormat,
): string {
  const value = {
    sparator: option?.sparator ?? " ",
    with_month: option?.with_month ?? true,
  };

  const currentDate = new Date(date);

  if (value.with_month) {
    return (
      currentDate.getDate() +
      value.sparator +
      Months[currentDate.getMonth()].substring(0, 3) +
      value.sparator +
      currentDate.getFullYear()
    );
  }

  return (
    currentDate.getDate() +
    value.sparator +
    (currentDate.getMonth() + 1 > 9
      ? (currentDate.getMonth() + 1).toString()
      : "0" + (currentDate.getMonth() + 1)) +
    value.sparator +
    currentDate.getFullYear()
  );
}

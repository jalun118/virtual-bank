export function CurrencyRupiahFormat(amount: number): string {
  const arrayData = new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
  }).formatToParts(amount);
  let text = "";

  arrayData.forEach((v) => {
    if (v.type !== "literal") {
      text += v.value;
    }
  });
  return text;
}

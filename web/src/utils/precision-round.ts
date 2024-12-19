export default function PrecisionRound(number: number, precision?: number) {
  const factor = Math.pow(10, precision ?? 12);
  return Math.round(number * factor) / factor;
}

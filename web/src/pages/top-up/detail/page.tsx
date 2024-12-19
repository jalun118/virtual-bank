import JsBarcode from "jsbarcode";
import { useCallback, useEffect, useRef } from "react";
import { CurrencyRupiahFormat } from "../../../utils/currency";
import { DateFormat } from "../../../utils/date";

export default function DetailTopUpPage() {
  const value = "4116503767565402";

  const containerRef = useRef<never>(null);

  const renderBarcode = useCallback(JsBarcode, [value, containerRef.current]);

  useEffect(() => {
    renderBarcode(containerRef.current, value, {
      displayValue: false,
      background: "transparent",
      lineColor: "black",
      margin: 0,
    });
  }, [renderBarcode, value]);

  return (
    <div className="container">
      <div className="m-auto max-w-md">
        <div className="rounded-2xl bg-gradient-to-t from-blue-500 via-indigo-500 via-50% to-cyan-500 p-10">
          <div className="rounded-xl bg-white p-7">
            <div className="flex flex-col items-center">
              <canvas ref={containerRef} className="w-full" />
              <p className="mb-4 mt-4 text-xl font-semibold tabular-nums">
                {value}
              </p>
            </div>
            <p className="border-t-2 border-dashed border-t-gray-400 pt-4">
              Total payment {CurrencyRupiahFormat(200000)}
            </p>
          </div>
          <p className="mt-4 text-center font-semibold text-white">
            Valid until {DateFormat(new Date(2025, 0, 2))}, 20:22
          </p>
        </div>
      </div>
    </div>
  );
}

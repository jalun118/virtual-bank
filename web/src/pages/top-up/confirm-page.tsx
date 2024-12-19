import { useNavigate } from "react-router-dom";
import { CurrencyRupiahFormat } from "../../utils/currency";
import PrecisionRound from "../../utils/precision-round";

interface iFormTopUp {
  method: string | null;
  admin_costs: number | null;
  amount: number | null;
}

export default function ConfirmPage({ formInfo }: { formInfo: iFormTopUp }) {
  const navigate = useNavigate();

  function handleSubmit() {
    navigate(`/top-up/detail?id=${btoa(JSON.stringify(formInfo))}`, {
      replace: true,
    });
  }

  return (
    <div className="mt-4">
      <div>
        <div className="border-b-2 border-dashed border-b-gray-400 pb-3">
          <div className="flex justify-between">
            <div className="mr-5">Amount</div>
            <div className="text-right tabular-nums">
              {CurrencyRupiahFormat(formInfo.amount ?? 0)}
            </div>
          </div>
          <div className="mt-3 flex justify-between">
            <div className="mr-5">Consts Admin</div>
            <div className="text-right tabular-nums">
              {CurrencyRupiahFormat(formInfo.admin_costs ?? 0)}
            </div>
          </div>
          <div className="mt-3 flex justify-between">
            <div className="mr-5">Method</div>
            <div className="text-right capitalize">{formInfo.method}</div>
          </div>
        </div>
        <div>
          <div className="mt-3 flex justify-between">
            <div className="mr-5">Total</div>
            <div className="text-right tabular-nums">
              {CurrencyRupiahFormat(
                PrecisionRound(
                  (formInfo.admin_costs ?? 0) + (formInfo.amount ?? 0),
                ),
              )}
            </div>
          </div>
        </div>
      </div>
      <div className="mt-12">
        <button
          onClick={() => handleSubmit()}
          type="submit"
          className="flex w-full justify-between rounded-lg border bg-gradient-to-t from-emerald-500 to-emerald-400 px-4 py-2 font-semibold text-white shadow-md active:bg-gradient-to-b active:shadow-inner"
        >
          <span>Top Up</span>
          <span className="text-right tabular-nums">
            {CurrencyRupiahFormat(
              PrecisionRound(
                (formInfo.admin_costs ?? 0) + (formInfo.amount ?? 0),
              ),
            )}
          </span>
        </button>
      </div>
    </div>
  );
}

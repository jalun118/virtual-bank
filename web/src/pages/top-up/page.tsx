import { useState } from "react";
import AmountPage from "./amount-page";
import ConfirmPage from "./confirm-page";
import MethodPage from "./method-page";

interface iFormTopUp {
  method: string | null;
  admin_costs: number | null;
  amount: number | null;
}

export default function TopUpPage() {
  const [getPage, SetPage] = useState(0);
  const [form, SetForm] = useState<iFormTopUp>({
    admin_costs: null,
    method: null,
    amount: null,
  });

  return (
    <div className="container">
      <div className="m-auto max-w-md">
        <h1 className="text-2xl font-semibold">Top Up</h1>
        <div className="mt-4">
          {getPage === 0 && (
            <MethodPage
              setForm={(v) =>
                SetForm((prev) => ({
                  ...prev,
                  admin_costs: v.admin_costs,
                  method: v.method,
                }))
              }
              setPage={(page) => SetPage(page)}
            />
          )}
          {getPage === 1 && (
            <AmountPage
              setAmount={(v) => SetForm((prev) => ({ ...prev, amount: v }))}
              handleNext={(v) => SetPage(v)}
            />
          )}
          {getPage === 2 && <ConfirmPage formInfo={form} />}
        </div>
      </div>
    </div>
  );
}

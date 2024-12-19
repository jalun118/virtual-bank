import { Link } from "react-router-dom";
import GetBadgeStatus from "../../components/GetBadgeStatus";
import { datas } from "../../dummy/type_transaction";
import { CurrencyRupiahFormat } from "../../utils/currency";

export default function TransactionPage() {
  return (
    <div className="container">
      <div className="m-auto max-w-md">
        <h1 className="text-2xl font-semibold">Transaction</h1>
        <div className="mt-5 grid gap-4">
          {datas.map((v) => (
            <Link to={"/transaction/" + v.id} key={v.id}>
              <div className="flex items-center justify-between rounded-xl border border-gray-400 bg-white p-4 shadow-sm dark:border-neutral-700 dark:bg-neutral-900 dark:text-neutral-400 md:p-5">
                <div className="flex items-center">
                  {v.amount < 0 ? (
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      strokeWidth={1.5}
                      stroke="currentColor"
                      className="mr-3 size-8"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M2.25 3h1.386c.51 0 .955.343 1.087.835l.383 1.437M7.5 14.25a3 3 0 0 0-3 3h15.75m-12.75-3h11.218c1.121-2.3 2.1-4.684 2.924-7.138a60.114 60.114 0 0 0-16.536-1.84M7.5 14.25 5.106 5.272M6 20.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Zm12.75 0a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Z"
                      />
                    </svg>
                  ) : (
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      strokeWidth={1.5}
                      stroke="currentColor"
                      className="mr-3 size-8"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M2.25 18.75a60.07 60.07 0 0 1 15.797 2.101c.727.198 1.453-.342 1.453-1.096V18.75M3.75 4.5v.75A.75.75 0 0 1 3 6h-.75m0 0v-.375c0-.621.504-1.125 1.125-1.125H20.25M2.25 6v9m18-10.5v.75c0 .414.336.75.75.75h.75m-1.5-1.5h.375c.621 0 1.125.504 1.125 1.125v9.75c0 .621-.504 1.125-1.125 1.125h-.375m1.5-1.5H21a.75.75 0 0 0-.75.75v.75m0 0H3.75m0 0h-.375a1.125 1.125 0 0 1-1.125-1.125V15m1.5 1.5v-.75A.75.75 0 0 0 3 15h-.75M15 10.5a3 3 0 1 1-6 0 3 3 0 0 1 6 0Zm3 0h.008v.008H18V10.5Zm-12 0h.008v.008H6V10.5Z"
                      />
                    </svg>
                  )}

                  <div className="flex flex-col justify-center">
                    <span>{v.amount < 0 ? "TRANFER" : "TOP UP"}</span>
                    <span className="text-sm">{v.dest_account_id}</span>
                  </div>
                </div>
                <div className="flex flex-col text-right">
                  {v.amount < 0 ? (
                    <span className="flex items-center tabular-nums text-red-600">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        strokeWidth={1.5}
                        stroke="currentColor"
                        className="size-4"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          d="M5 12h14"
                        />
                      </svg>
                      {CurrencyRupiahFormat(Math.abs(v.amount))}
                    </span>
                  ) : (
                    <span className="flex items-center tabular-nums text-green-600">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        strokeWidth={1.5}
                        stroke="currentColor"
                        className="size-4"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          d="M12 4.5v15m7.5-7.5h-15"
                        />
                      </svg>
                      {CurrencyRupiahFormat(Math.abs(v.amount))}
                    </span>
                  )}
                  <span className="flex justify-end text-sm">
                    {GetBadgeStatus(v.status)}
                  </span>
                </div>
              </div>
            </Link>
          ))}
        </div>
      </div>
    </div>
  );
}

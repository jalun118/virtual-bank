import { useParams } from "react-router-dom";
import GetBadgeStatus from "../../../components/GetBadgeStatus";
import NotFoundIcon from "../../../components/NotFoundIcon";
import { datas } from "../../../dummy/type_transaction";
import { CurrencyRupiahFormat } from "../../../utils/currency";
import { DateFormat } from "../../../utils/date";

function CreateSerial(dest_account: string, date: string | Date): string {
  const cur = new Date(
    new Date(date).getUTCMilliseconds() + parseInt(dest_account),
  );
  const dateStruct = [
    cur.getFullYear().toString(),
    cur.getMonth() + 1 > 9
      ? (cur.getMonth() + 1).toString()
      : "0" + (cur.getMonth() + 1),
    cur.getDate() > 9 ? cur.getDate().toString() : "0" + cur.getDate(),
    cur.getHours() > 9 ? cur.getHours().toString() : "0" + cur.getHours(),
    cur.getMinutes() > 9 ? cur.getMinutes().toString() : "0" + cur.getMinutes(),
    cur.getSeconds() > 9 ? cur.getSeconds().toString() : "0" + cur.getSeconds(),
  ];
  return dateStruct.join("");
}

export default function TransactionDetailPage() {
  const { id } = useParams();
  const selectData = datas.find((v) => v.id === id);

  if (!selectData) {
    return (
      <div className="container mt-5">
        <NotFoundIcon />
      </div>
    );
  }

  return (
    <div className="container mt-5">
      <div className="mx-auto max-w-md rounded-xl border border-gray-400 p-5 text-base">
        <div className="flex flex-col gap-y-1.5">
          <div className="flex justify-between">
            <div className="mr-5">Status</div>
            <div className="truncate whitespace-nowrap text-right">
              {GetBadgeStatus(selectData?.status)}
            </div>
          </div>
          <div className="flex justify-between">
            <div className="mr-5">Destination</div>
            <div className="flex flex-col items-end">
              <div className="flex items-center text-right">
                <div className="me-2 truncate whitespace-nowrap">
                  {selectData?.dest_account_id}
                </div>
                <button className="rounded-md bg-gray-200 p-0.5 active:bg-gray-300">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className="size-5"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M15.75 17.25v3.375c0 .621-.504 1.125-1.125 1.125h-9.75a1.125 1.125 0 0 1-1.125-1.125V7.875c0-.621.504-1.125 1.125-1.125H6.75a9.06 9.06 0 0 1 1.5.124m7.5 10.376h3.375c.621 0 1.125-.504 1.125-1.125V11.25c0-4.46-3.243-8.161-7.5-8.876a9.06 9.06 0 0 0-1.5-.124H9.375c-.621 0-1.125.504-1.125 1.125v3.5m7.5 10.375H9.375a1.125 1.125 0 0 1-1.125-1.125v-9.25m12 6.625v-1.875a3.375 3.375 0 0 0-3.375-3.375h-1.5a1.125 1.125 0 0 1-1.125-1.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H9.75"
                    />
                  </svg>
                </button>
              </div>
              <div className="text-right font-bold uppercase">Kuda Jemping</div>
            </div>
          </div>

          <div className="flex justify-between">
            <div className="mr-5">ID Transaction</div>
            <div className="relative flex items-center">
              <div className="me-5 max-w-36 truncate whitespace-nowrap text-right">
                {selectData?.id}
              </div>
              <button className="absolute right-0 rounded-md bg-gray-200 p-0.5 active:bg-gray-300">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  strokeWidth={1.5}
                  stroke="currentColor"
                  className="size-5"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    d="M15.75 17.25v3.375c0 .621-.504 1.125-1.125 1.125h-9.75a1.125 1.125 0 0 1-1.125-1.125V7.875c0-.621.504-1.125 1.125-1.125H6.75a9.06 9.06 0 0 1 1.5.124m7.5 10.376h3.375c.621 0 1.125-.504 1.125-1.125V11.25c0-4.46-3.243-8.161-7.5-8.876a9.06 9.06 0 0 0-1.5-.124H9.375c-.621 0-1.125.504-1.125 1.125v3.5m7.5 10.375H9.375a1.125 1.125 0 0 1-1.125-1.125v-9.25m12 6.625v-1.875a3.375 3.375 0 0 0-3.375-3.375h-1.5a1.125 1.125 0 0 1-1.125-1.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H9.75"
                  />
                </svg>
              </button>
            </div>
          </div>
          <div className="flex justify-between">
            <div className="mr-5">Time</div>
            <div className="truncate whitespace-nowrap text-right">
              {new Date(selectData?.transaction_date ?? "").getHours()}:
              {new Date(selectData?.transaction_date ?? "").getMinutes()}
            </div>
          </div>
          <div className="flex justify-between">
            <div className="mr-5">Date</div>
            <div className="truncate whitespace-nowrap text-right">
              {DateFormat(selectData?.transaction_date ?? "")}
            </div>
          </div>
          <div className="flex justify-between">
            <div className="mr-5">Serial</div>

            <div className="relative flex items-center">
              <div className="me-2 truncate whitespace-nowrap text-right">
                {CreateSerial(
                  selectData?.dest_account_id ?? "",
                  selectData?.transaction_date ?? "",
                )}
              </div>
              <button className="rounded-md bg-gray-200 p-0.5 active:bg-gray-300">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  strokeWidth={1.5}
                  stroke="currentColor"
                  className="size-5"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    d="M15.75 17.25v3.375c0 .621-.504 1.125-1.125 1.125h-9.75a1.125 1.125 0 0 1-1.125-1.125V7.875c0-.621.504-1.125 1.125-1.125H6.75a9.06 9.06 0 0 1 1.5.124m7.5 10.376h3.375c.621 0 1.125-.504 1.125-1.125V11.25c0-4.46-3.243-8.161-7.5-8.876a9.06 9.06 0 0 0-1.5-.124H9.375c-.621 0-1.125.504-1.125 1.125v3.5m7.5 10.375H9.375a1.125 1.125 0 0 1-1.125-1.125v-9.25m12 6.625v-1.875a3.375 3.375 0 0 0-3.375-3.375h-1.5a1.125 1.125 0 0 1-1.125-1.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H9.75"
                  />
                </svg>
              </button>
            </div>
          </div>
          {(selectData?.description ?? "") !== "" && (
            <div className="flex justify-between">
              <div className="mr-5">Description</div>
              <div className="text-right">{selectData?.description}</div>
            </div>
          )}
        </div>
        <div className="my-3 flex flex-col gap-y-1 border-y-2 border-dashed border-y-gray-700 py-2">
          <div className="flex justify-between">
            <div className="mr-5">Amount</div>
            <div className="truncate whitespace-nowrap text-right">
              {CurrencyRupiahFormat(selectData?.amount ?? 0)}
            </div>
          </div>
        </div>
        <div>
          <div className="flex justify-between">
            <div className="mr-5">Total</div>
            <div className="truncate whitespace-nowrap text-right">
              {CurrencyRupiahFormat(selectData?.amount ?? 0)}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

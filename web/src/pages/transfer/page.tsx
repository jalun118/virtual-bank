import { useState } from "react";
import AmountTransfer from "./amount-transfer";
import DetailAccount from "./detail-account";
import PinPage from "./pin-page";
import ProccessTransfer from "./proccess-transfer";
import VerifyPage from "./verify-page";

export default function TransferPage() {
  const [getPage, setPage] = useState(0);

  return (
    <div className="container">
      <div className="m-auto max-w-md">
        <h1 className="text-xl font-semibold">Transfer</h1>
        <div className="mt-5 px-1">
          {getPage === 0 && <VerifyPage setPage={(v) => setPage(v)} />}
          {getPage === 1 && <DetailAccount setPage={(v) => setPage(v)} />}
          {getPage === 2 && <AmountTransfer setPage={(v) => setPage(v)} />}
          {getPage === 3 && <PinPage setPage={(v) => setPage(v)} />}
          {getPage === 4 && <ProccessTransfer />}
        </div>
      </div>
    </div>
  );
}

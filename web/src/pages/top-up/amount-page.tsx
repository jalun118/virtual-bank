import { useFormik } from "formik";
import { number as yupNumber, object as yupObject } from "yup";

interface iDefaultValue {
  amount: number;
}

const defaultValue: iDefaultValue = {
  amount: 0,
};

const topUpSchema = yupObject({
  amount: yupNumber()
    .required("Required")
    .min(10000, "Minimum top-up Rp10.000"),
});

export default function AmountPage({
  setAmount,
  handleNext,
}: {
  setAmount: (amount: number) => void;
  handleNext: (nextPage: number) => void;
}) {
  const { values, touched, handleChange, handleSubmit, errors } = useFormik({
    initialValues: defaultValue,
    onSubmit: (values) => {
      setAmount(values.amount);
      handleNext(2);
    },
    validationSchema: topUpSchema,
  });

  return (
    <div>
      <h2 className="text-lg font-medium">Amount</h2>
      <form onSubmit={handleSubmit}>
        <div className="mt-2">
          <input
            type="number"
            autoComplete="off"
            autoCorrect="off"
            aria-autocomplete="none"
            spellCheck="false"
            name="amount"
            className="block w-full rounded-lg px-4 py-3 text-black"
            placeholder="Rp 0"
            value={values.amount}
            onChange={handleChange}
          />
          {touched.amount && errors.amount && (
            <div className="text-sm font-semibold text-red-500">
              {errors.amount}
            </div>
          )}
        </div>
        <div className="mt-5">
          <button
            type="submit"
            className="flex justify-between rounded-lg border bg-gradient-to-t from-cyan-500 to-cyan-300 px-5 py-2 font-semibold text-white shadow-md active:bg-gradient-to-b active:shadow-inner"
          >
            Confirm
          </button>
        </div>
      </form>
    </div>
  );
}

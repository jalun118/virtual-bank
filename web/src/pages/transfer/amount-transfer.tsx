import { useFormik } from "formik";
import {
  number as yupNumber,
  object as yupObject,
  string as yupString,
} from "yup";

interface iDefaultValue {
  amount: number;
  description: string;
}

const defaultValue: iDefaultValue = {
  amount: 0,
  description: "",
};

const transferSchema = yupObject({
  amount: yupNumber()
    .required("Required")
    .min(10000, "Minimum transfer Rp10.000"),
  description: yupString().max(3000, "Maximum 3000 character"),
});

export default function AmountTransfer({
  setPage,
}: {
  setPage: (page: number) => void;
}) {
  const { touched, errors, values, handleChange, handleSubmit } = useFormik({
    initialValues: defaultValue,
    onSubmit: (values) => {
      console.log(values);
      setPage(3);
    },
    validationSchema: transferSchema,
  });

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="amount" className="mb-2 block font-semibold">
            Amount
          </label>
          <input
            id="amount"
            name="amount"
            value={values.amount}
            onChange={handleChange}
            autoComplete="off"
            autoCorrect="off"
            aria-autocomplete="none"
            spellCheck="false"
            className="block w-full rounded-md px-3 py-2 text-black"
            placeholder="200000"
          />
          {touched.amount && errors.amount && (
            <div className="text-sm font-semibold text-red-500">
              {errors.amount}
            </div>
          )}
        </div>
        <div className="mt-4">
          <label htmlFor="description" className="mb-2 block font-semibold">
            Description
          </label>
          <textarea
            id="description"
            name="description"
            value={values.description}
            onChange={handleChange}
            autoComplete="off"
            autoCorrect="off"
            aria-autocomplete="none"
            spellCheck="false"
            className="block h-40 w-full resize-none rounded-md px-3 py-2 text-black"
          />
          {touched.description && errors.description && (
            <div className="text-sm font-semibold text-red-500">
              {errors.description}
            </div>
          )}
        </div>
        <div className="mt-7">
          <button
            type="submit"
            className="flex justify-between rounded-lg border bg-gradient-to-t from-cyan-500 to-cyan-300 px-3 py-2 font-semibold text-white shadow-md active:bg-gradient-to-b active:shadow-inner"
          >
            Continue
          </button>
        </div>
      </form>
    </div>
  );
}

import { useFormik } from "formik";
import { object as yupObject, string as yupString } from "yup";

interface iDefaultValue {
  username: string;
}

const editUsernameSchema = yupObject({
  username: yupString().required("Required"),
});

export default function EditInformationPage() {
  const initValue: iDefaultValue = {
    username: "",
  };

  const { touched, errors, values, handleChange, handleSubmit } = useFormik({
    initialValues: initValue,
    onSubmit: (value) => {
      console.log(value);
    },
    validationSchema: editUsernameSchema,
  });

  return (
    <div className="container">
      <div className="m-auto max-w-md">
        <h1 className="text-xl font-semibold">Edit Username</h1>
        <div className="mt-5">
          <form onSubmit={handleSubmit}>
            <div className="max-w-sm">
              <input
                type="text"
                autoComplete="off"
                autoCorrect="off"
                aria-autocomplete="none"
                spellCheck="false"
                className="block w-full rounded-lg px-4 py-3 text-black"
                placeholder="kuda_jemping"
                value={values.username}
                onChange={handleChange}
              />
              {touched.username && errors.username && (
                <div className="mt-2 text-sm font-semibold text-red-500">
                  {errors.username}
                </div>
              )}
            </div>
            <div className="mt-5">
              <button
                type="submit"
                className="rounded-md bg-blue-600 px-4 py-2 font-semibold text-white active:bg-blue-700"
              >
                Save
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}

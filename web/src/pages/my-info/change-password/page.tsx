import { useFormik } from "formik";
import { useState } from "react";
import { Link } from "react-router-dom";
import { object as yupObject, string as yupString } from "yup";

interface iDefaultValue {
  old_password: string;
  new_password: string;
  confirm_password: string;
}

const changePasswordSchema = yupObject({
  confirm_password: yupString()
    .required("Required")
    .min(5, "Minimum 5 characters"),
  new_password: yupString().required("Required").min(5, "Minimum 5 characters"),
  old_password: yupString().required("Required").min(5, "Minimum 5 characters"),
});

export default function ChangePasswordPage() {
  const [isShowFields, setHideFields] = useState({
    old_password: false,
    new_password: false,
    confirm_password: false,
  });

  function HandleToggleShow(field: keyof iDefaultValue) {
    setHideFields((prev) => ({
      ...prev,
      [field]: !prev[field],
    }));
  }

  const initValue: iDefaultValue = {
    confirm_password: "",
    new_password: "",
    old_password: "",
  };

  const { touched, errors, handleSubmit, handleChange, values } = useFormik({
    initialValues: initValue,
    onSubmit: (value, action) => {
      console.log(value);
      if (value.confirm_password !== value.new_password) {
        action.setErrors({ confirm_password: "Password not match" });
        return;
      }

      location.replace("/");
    },
    validationSchema: changePasswordSchema,
  });

  return (
    <div className="container">
      <div className="m-auto max-w-md">
        <h1 className="text-xl font-semibold">Change Password</h1>
        <div className="mt-3">
          <form onSubmit={handleSubmit}>
            <div>
              <label
                htmlFor="old_password"
                className="mb-2 block font-semibold"
              >
                Old Password
              </label>
              <div className="relative">
                <input
                  type={isShowFields.old_password ? "text" : "password"}
                  id="old_password"
                  name="old_password"
                  autoComplete="off"
                  autoCorrect="off"
                  aria-autocomplete="none"
                  spellCheck="false"
                  className="block w-full rounded-md px-3 py-2 pe-9 text-black"
                  value={values.old_password}
                  onChange={handleChange}
                />
                <button
                  onClick={() => HandleToggleShow("old_password")}
                  className="absolute inset-y-0 right-2"
                >
                  {isShowFields.old_password ? (
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      strokeWidth={1.5}
                      stroke="currentColor"
                      className="size-6"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M3.98 8.223A10.477 10.477 0 0 0 1.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.451 10.451 0 0 1 12 4.5c4.756 0 8.773 3.162 10.065 7.498a10.522 10.522 0 0 1-4.293 5.774M6.228 6.228 3 3m3.228 3.228 3.65 3.65m7.894 7.894L21 21m-3.228-3.228-3.65-3.65m0 0a3 3 0 1 0-4.243-4.243m4.242 4.242L9.88 9.88"
                      />
                    </svg>
                  ) : (
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      strokeWidth={1.5}
                      stroke="currentColor"
                      className="size-6"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M2.036 12.322a1.012 1.012 0 0 1 0-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178Z"
                      />
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z"
                      />
                    </svg>
                  )}
                </button>
              </div>
              {touched.old_password && errors.old_password && (
                <div className="text-sm font-semibold text-red-500">
                  {errors.old_password}
                </div>
              )}
            </div>
            <div className="mt-2">
              <label
                htmlFor="new_password"
                className="mb-2 block font-semibold"
              >
                New Password
              </label>
              <div className="relative">
                <input
                  type={isShowFields.new_password ? "text" : "password"}
                  id="new_password"
                  name="new_password"
                  autoComplete="off"
                  autoCorrect="off"
                  aria-autocomplete="none"
                  spellCheck="false"
                  className="block w-full rounded-md px-3 py-2 text-black"
                  value={values.new_password}
                  onChange={handleChange}
                />
                <button
                  onClick={() => HandleToggleShow("new_password")}
                  className="absolute inset-y-0 right-2"
                >
                  {isShowFields.new_password ? (
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      strokeWidth={1.5}
                      stroke="currentColor"
                      className="size-6"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M3.98 8.223A10.477 10.477 0 0 0 1.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.451 10.451 0 0 1 12 4.5c4.756 0 8.773 3.162 10.065 7.498a10.522 10.522 0 0 1-4.293 5.774M6.228 6.228 3 3m3.228 3.228 3.65 3.65m7.894 7.894L21 21m-3.228-3.228-3.65-3.65m0 0a3 3 0 1 0-4.243-4.243m4.242 4.242L9.88 9.88"
                      />
                    </svg>
                  ) : (
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      strokeWidth={1.5}
                      stroke="currentColor"
                      className="size-6"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M2.036 12.322a1.012 1.012 0 0 1 0-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178Z"
                      />
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z"
                      />
                    </svg>
                  )}
                </button>
              </div>
              {touched.new_password && errors.new_password && (
                <div className="text-sm font-semibold text-red-500">
                  {errors.new_password}
                </div>
              )}
            </div>
            <div className="mt-2">
              <label
                htmlFor="confirm_password"
                className="mb-2 block font-semibold"
              >
                Confirm Password
              </label>
              <div className="relative">
                <input
                  type={isShowFields.confirm_password ? "text" : "password"}
                  id="confirm_password"
                  name="confirm_password"
                  autoComplete="off"
                  autoCorrect="off"
                  aria-autocomplete="none"
                  spellCheck="false"
                  className="block w-full rounded-md px-3 py-2 text-black"
                  value={values.confirm_password}
                  onChange={handleChange}
                />
                <button
                  onClick={() => HandleToggleShow("confirm_password")}
                  className="absolute inset-y-0 right-2"
                >
                  {isShowFields.confirm_password ? (
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      strokeWidth={1.5}
                      stroke="currentColor"
                      className="size-6"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M3.98 8.223A10.477 10.477 0 0 0 1.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.451 10.451 0 0 1 12 4.5c4.756 0 8.773 3.162 10.065 7.498a10.522 10.522 0 0 1-4.293 5.774M6.228 6.228 3 3m3.228 3.228 3.65 3.65m7.894 7.894L21 21m-3.228-3.228-3.65-3.65m0 0a3 3 0 1 0-4.243-4.243m4.242 4.242L9.88 9.88"
                      />
                    </svg>
                  ) : (
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      strokeWidth={1.5}
                      stroke="currentColor"
                      className="size-6"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M2.036 12.322a1.012 1.012 0 0 1 0-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178Z"
                      />
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z"
                      />
                    </svg>
                  )}
                </button>
              </div>
              {touched.confirm_password && errors.confirm_password && (
                <div className="text-sm font-semibold text-red-500">
                  {errors.confirm_password}
                </div>
              )}
            </div>
            <div className="mt-2">
              <p>
                Do you need help?{" "}
                <Link to="/" className="font-semibold underline">
                  Lost Password
                </Link>
              </p>
            </div>
            <div className="mt-2">
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

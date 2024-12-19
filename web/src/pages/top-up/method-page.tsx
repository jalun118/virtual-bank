interface dataForm {
  method: string;
  admin_costs: number;
}

interface methodProvider {
  name: string;
  admin_costs: number;
  image_url: string;
}

const listMethodProvider: methodProvider[] = [
  {
    admin_costs: 1500,
    image_url: "https://www.indomaret.co.id/Assets/image/logo.png",
    name: "Indomart",
  },
  {
    admin_costs: 1000,
    image_url:
      "https://upload.wikimedia.org/wikipedia/commons/8/86/Alfamart_logo.svg",
    name: "Alfamart",
  },
];

export default function MethodPage({
  setPage,
  setForm,
}: {
  setPage: (page: number) => void;
  setForm: (data: dataForm) => void;
}) {
  function hadleSetMethod(data: dataForm) {
    setPage(1);
    setForm(data);
  }

  return (
    <div>
      <h2 className="text-lg font-medium">Method TopUp</h2>
      <div className="mt-3 flex gap-4">
        {listMethodProvider.map((provider, idx) => (
          <button
            onClick={() =>
              hadleSetMethod({
                admin_costs: provider.admin_costs,
                method: provider.name.toLowerCase(),
              })
            }
            key={idx}
            className="flex flex-col items-center rounded-xl border border-gray-400 px-1.5 py-2 hover:bg-gray-100"
          >
            <span className="flex aspect-square w-14 items-center justify-center">
              <img
                src={provider.image_url}
                alt={provider.name}
                className="h-auto w-full"
              />
            </span>
            <span className="mt-1.5 font-semibold">{provider.name}</span>
          </button>
        ))}
      </div>
    </div>
  );
}

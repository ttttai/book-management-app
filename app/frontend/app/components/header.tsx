type HeaderProps = {
  title: string;
};

const Header = ({ title }: HeaderProps) => {
  return (
    <header className="bg-black text-white py-4 px-6 shadow-md">
      <h1 className="text-xl font-bold">{title}</h1>
    </header>
  );
};

export default Header;

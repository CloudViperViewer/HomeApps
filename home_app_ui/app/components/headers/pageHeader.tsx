interface PageHeaderProps {
  headerText: string;
}

function PageHeader({ headerText }: PageHeaderProps) {
  return (
    <>
      <h1 className="text-2xl font-bold">{headerText}</h1>
      <hr className="mt-2" />
    </>
  );
}

export default PageHeader;

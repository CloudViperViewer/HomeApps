interface pageHeaderProps {
  headerText: string;
}

function pageHeader({ headerText }: pageHeaderProps) {
  return (
    <>
      <h1 className="text-2xl">
        <b>{headerText}</b>
      </h1>
      <hr className="mt-2" />
    </>
  );
}

export default pageHeader;

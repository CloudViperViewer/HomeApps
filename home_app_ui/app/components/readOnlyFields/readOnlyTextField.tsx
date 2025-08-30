interface readOnlyTextFieldProps {
  label: string;
  value: string | number | boolean;
  className?: string;
}

function readOnlyTextField({
  label,
  value,
  className = "",
}: readOnlyTextFieldProps) {
  return (
    <div className={`mt-1, mb-1 ${className}`}>
      <div>
        <b>{label}</b>
      </div>
      <div>{value}</div>
    </div>
  );
}

export default readOnlyTextField;

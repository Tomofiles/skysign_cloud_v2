import React, { useEffect, useState } from "react";

const dispFloor = (num, digit) => {
  return Math.floor(num * Math.pow(10, digit) ) / Math.pow(10, digit);
}

const NumberValue = props => {
  const [ value, setValue ] = useState("-");

  useEffect(() => {
    if (props.value === "-") return;
    setValue(dispFloor(props.value, 8));
  }, [ props.value ])

  return (
    <span>{value}</span>
  );
}

export default NumberValue; 
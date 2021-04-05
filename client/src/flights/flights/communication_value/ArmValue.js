import React, { useEffect, useState } from "react";

const ArmValue = props => {
  const [ value, setValue ] = useState("-");

  useEffect(() => {
    if (props.value === "-") return;
    if (props.value) {
      setValue("ARMED");
    } else {
      setValue("DISARMED");
    }
  }, [ props.value ])

  return (
    <span>{value}</span>
  );
}

export default ArmValue; 
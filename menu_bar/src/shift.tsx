import { useEffect, useState } from "react";
import { getPreferenceValues, MenuBarExtra } from "@raycast/api";
import fs from "fs/promises";

export default function Command() {
  const [shift, setShift] = useState<boolean>();
  const { file } = getPreferenceValues();

  useEffect(() => {
    if (!file) return setShift(false);
    const interval = setInterval(async () => {
      const text = await fs.readFile(file, "utf-8");
      setShift(text === "1");
      console.log("hello")
    }, 100);

    return () => clearInterval(interval);
  }, []);

  return (
    <MenuBarExtra icon={{ source: shift ? "filled.svg" : "outlined.svg" }} isLoading title={""}>
      <MenuBarExtra.Item title="DO NOT OPEN THIS. YOU ARE BAD" />
    </MenuBarExtra>
  );
}

import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const ProgrammableTokenTransferModule = buildModule("ProgrammableTokenTransferModule", (m) => {
  const transfer = m.contract("ProgrammableTokenTransfer", [m.getParameter("owner"), m.getParameter("router")]);

  return { transfer };
});

export default ProgrammableTokenTransferModule;

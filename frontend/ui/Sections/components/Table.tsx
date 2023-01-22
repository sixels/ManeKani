import {
  Table,
  TableCellProps,
  TableContainer,
  TableProps,
  Tbody,
  Td,
  Tr,
} from "@chakra-ui/react";

export type HeadlessTableProps = {
  table: {
    title: string;
    value: string | JSX.Element;
    valueProps?: TableCellProps;
    titleProps?: TableCellProps;
    rowProps?: any;
  }[];
  props?: TableProps;
};

export function HeadlessTable({ table, props }: HeadlessTableProps) {
  const items = table.map(
    ({ title, value, valueProps, titleProps, rowProps }, i) => {
      return (
        <Tr key={i} {...rowProps}>
          <Td
            textTransform={"uppercase"}
            fontWeight={"bold"}
            color="GrayText"
            fontSize="md"
            {...titleProps}
          >
            {title}
          </Td>
          <Td
            w="full"
            justifyItems={"end"}
            {...valueProps}
            whiteSpace="normal"
            fontSize="md"
          >
            {value}
          </Td>
        </Tr>
      );
    }
  );

  return (
    <TableContainer py={2}>
      <Table variant={"unstyled"} size="sm" {...props}>
        <Tbody>{items}</Tbody>
      </Table>
    </TableContainer>
  );
}

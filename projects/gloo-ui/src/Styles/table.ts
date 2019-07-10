import styled from '@emotion/styled/macro';
import { colors } from 'Styles/colors';

export const TableActionCircle = styled.div`
  display: inline-block;
  width: 18px;
  height: 18px;
  line-height: 18px;
  text-align: center;
  color: ${colors.novemberGrey};
  border-radius: 18px;
  cursor: pointer;

  background: ${colors.marchGrey};

  &:hover,
  &:focus {
    background: ${colors.mayGrey};
  }

  &:active {
    background: ${colors.marchGrey};
  }
`;

export const TableHealthCircleHolder = styled.div`
  display: inline;

  > div {
    width: 10px;
    height: 10px;
    margin-left: 0;
    margin-right: 5px;
  }
`;

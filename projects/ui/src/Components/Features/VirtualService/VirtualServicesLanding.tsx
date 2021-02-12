import React, { useState } from 'react';
import styled from '@emotion/styled/macro';
import {
  VirtualServicesTable,
  VirtualServicesPageTable,
} from './VirtualServicesTable';
import { SoloInput } from 'Components/Common/SoloInput';
import { SoloRadioGroup } from 'Components/Common/SoloRadioGroup';
import { VirtualServiceStatus } from 'proto/github.com/solo-io/solo-apis/api/gloo/gateway/v1/virtual_service_pb';
import { colors } from 'Styles/colors';

const VirtualServiceLandingContainer = styled.div`
  display: grid;
  grid-template-columns: 200px 1fr;
  grid-gap: 28px;
`;

const HorizontalDivider = styled.div`
  position: relative;
  height: 1px;
  width: 100%;
  background: ${colors.marchGrey};
  margin: 35px 0;

  div {
    position: absolute;
    display: block;
    left: 0;
    right: 0;
    top: 50%;
    margin: -9px auto 0;
    width: 105px;
    text-align: center;
    color: ${colors.septemberGrey};
    background: ${colors.januaryGrey};
  }
`;

export const VirtualServicesLanding = () => {
  const [nameFilter, setNameFilter] = useState('');
  const [statusFilter, setStatusFilter] = useState<
    VirtualServiceStatus.StateMap[keyof VirtualServiceStatus.StateMap]
  >();

  const changeNameFilter = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNameFilter(e.target.value);
  };

  const changeStatus = (idSelected: string | number | undefined) => {
    setStatusFilter(
      idSelected as VirtualServiceStatus.StateMap[keyof VirtualServiceStatus.StateMap]
    );
  };

  return (
    <VirtualServiceLandingContainer>
      <div>
        <SoloInput
          value={nameFilter}
          onChange={changeNameFilter}
          placeholder={'Filter by name...'}
        />

        <HorizontalDivider>
          <div>Status Filter</div>
        </HorizontalDivider>

        <SoloRadioGroup
          options={[
            {
              displayName: 'Accepted',
              id: VirtualServiceStatus.State.ACCEPTED,
            },
            {
              displayName: 'Rejected',
              id: VirtualServiceStatus.State.REJECTED,
            },
            {
              displayName: 'Pending',
              id: VirtualServiceStatus.State.PENDING,
            },

            {
              displayName: 'Warning',
              id: VirtualServiceStatus.State.WARNING,
            },
          ]}
          currentSelection={statusFilter}
          onChange={changeStatus}
        />
      </div>
      <VirtualServicesPageTable
        nameFilter={nameFilter}
        statusFilter={statusFilter}
      />
    </VirtualServiceLandingContainer>
  );
};

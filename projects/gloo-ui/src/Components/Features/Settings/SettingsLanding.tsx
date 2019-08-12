import * as React from 'react';
/** @jsx jsx */
import { jsx } from '@emotion/core';

import styled from '@emotion/styled/macro';
import { RouteComponentProps, Route, Switch, Redirect } from 'react-router';
import { colors } from 'Styles';
import {
  ListingFilter,
  TypeFilterProps,
  StringFilterProps,
  CheckboxFilterProps,
  RadioFilterProps
} from 'Components/Common/ListingFilter';
import { SecretsPage } from './SecretsPage';
import { WatchedNamespacesPage } from './WatchedNamespacesPage';
import { SecurityPage } from './SecurityPage';
import { Breadcrumb } from 'Components/Common/Breadcrumb';
import {
  Secret,
  AwsSecret,
  AzureSecret,
  TlsSecret
} from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/secret_pb';
import { useListSecrets, useCreateSecret, useDeleteSecret } from 'Api';
import {
  ListSecretsRequest,
  CreateSecretRequest,
  DeleteSecretRequest
} from 'proto/github.com/solo-io/solo-projects/projects/grpcserver/api/v1/secret_pb';
import { NamespacesContext } from 'GlooIApp';
import { SuccessModal } from 'Components/Common/SuccessModal';
import { SecretValuesType } from './SecretForm';
import { ResourceRef } from 'proto/github.com/solo-io/solo-kit/api/v1/ref_pb';
import { useGetSecretsListV2 } from 'Api/v2/useSecretClientV2';
import { secrets } from 'Api/v2/SecretClient';

const PageChoiceFilter: TypeFilterProps = {
  id: 'pageChoice',
  options: [
    {
      displayName: 'Security'
    },
    {
      displayName: 'Watched Namespaces'
    },
    {
      displayName: 'Secrets'
    }
  ],
  choice: 'Security'
};

const Heading = styled.div`
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
`;

interface Props extends RouteComponentProps {}

export const SettingsLanding = (props: Props) => {
  const [awsSecrets, setAwsSecrets] = React.useState<Secret.AsObject[]>([]);
  const [azureSecrets, setAzureSecrets] = React.useState<Secret.AsObject[]>([]);
  const [tlsSecrets, setTlsSecrets] = React.useState<Secret.AsObject[]>([]);
  const [oAuthSecrets, setOAuthSecrets] = React.useState<Secret.AsObject[]>([]);
  const namespaces = React.useContext(NamespacesContext);
  const [showSuccessModal, setShowSuccessModal] = React.useState(false);

  const { data, loading, error, setNewVariables } = useGetSecretsListV2({
    namespaces: namespaces.namespacesList
  });
  const [allSecrets, setAllSecrets] = React.useState<Secret.AsObject[]>([]);
  React.useEffect(() => {
    if (!!data) {
      setAllSecrets(data.toObject().secretsList);
    }
    return () => {
      setShowSuccessModal(false);
    };
  }, [data, showSuccessModal]);

  React.useEffect(() => {
    if (data && allSecrets) {
      setAwsSecrets(allSecrets.filter(s => !!s.aws));
      setAzureSecrets(allSecrets.filter(s => !!s.azure));
      setOAuthSecrets(allSecrets.filter(s => !!s.extension));
      setTlsSecrets(allSecrets.filter(s => !!s.tls));
    }
  }, [allSecrets.length, showSuccessModal]);

  if (!data || (!data && loading)) {
    return <div>Loading...</div>;
  }
  const locationEnding = props.location.pathname.split('/settings/')[1];

  const startingChoice =
    locationEnding && locationEnding.length
      ? locationEnding === 'namespaces'
        ? 'Watched Namespaces'
        : locationEnding.charAt(0).toUpperCase() + locationEnding.slice(1)
      : 'Security';

  const pageChanged = (
    strings: StringFilterProps[],
    types: TypeFilterProps[],
    checkboxes: CheckboxFilterProps[],
    radios: RadioFilterProps[]
  ) => {
    const newChoice = types.find(type => type.id === 'pageChoice')!.choice!;
    const newPageLocation =
      newChoice === 'Watched Namespaces'
        ? 'namespaces'
        : newChoice.toLowerCase();

    props.history.push({
      pathname: `/settings/${newPageLocation}`
    });
  };

  const listDisplay = (
    strings: StringFilterProps[],
    types: TypeFilterProps[],
    checkboxes: CheckboxFilterProps[],
    radios: RadioFilterProps[]
  ): React.ReactNode => {
    return (
      <React.Fragment>
        <Switch>
          <Route
            path='/settings/security/'
            render={() => (
              <SecurityPage
                tlsSecrets={tlsSecrets}
                oAuthSecrets={oAuthSecrets}
                onCreateSecret={createSecret}
                onDeleteSecret={deleteSecret}
              />
            )}
          />
          <Route
            path='/settings/namespaces/'
            render={() => <WatchedNamespacesPage />}
          />
          <Route
            path='/settings/secrets/'
            render={() => (
              <SecretsPage
                awsSecrets={awsSecrets}
                azureSecrets={azureSecrets}
                onCreateSecret={createSecret}
                onDeleteSecret={deleteSecret}
              />
            )}
          />

          <Redirect exact from='/settings/' to='/settings/security/' />
        </Switch>
      </React.Fragment>
    );
  };

  async function createSecret(
    values: SecretValuesType,
    secretKind: Secret.KindCase
  ) {
    const {
      secretResourceRef: { name, namespace }
    } = values;
    try {
      await secrets.createSecret({ name, namespace, values, secretKind });
    } catch (error) {
      // TODO: show error modal
      console.error('error', error);
    }
    setNewVariables({ namespaces: namespaces.namespacesList });
    setShowSuccessModal(true);
  }

  async function deleteSecret(
    name: string,
    namespace: string,
    secretKind: Secret.KindCase
  ) {
    if (secretKind === Secret.KindCase.AWS) {
      setAwsSecrets(awsSecrets =>
        awsSecrets.filter(s => s.metadata!.name !== name)
      );
    }
    if (secretKind === Secret.KindCase.AZURE) {
      setAzureSecrets(azureSecrets =>
        azureSecrets.filter(s => s.metadata!.name !== name)
      );
    }
    if (secretKind === Secret.KindCase.TLS) {
      setTlsSecrets(tlsSecrets =>
        tlsSecrets.filter(s => s.metadata!.name !== name)
      );
    }
    if (secretKind === Secret.KindCase.EXTENSION) {
      setOAuthSecrets(oAuthSecrets =>
        oAuthSecrets.filter(s => s.metadata!.name !== name)
      );
    }
    try {
      await secrets.deleteSecret({ name, namespace });
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <div>
      <Heading>
        <Breadcrumb />
      </Heading>
      <SuccessModal
        visible={showSuccessModal}
        successMessage='Secret added successfully'
      />
      <ListingFilter
        types={[{ ...PageChoiceFilter, choice: startingChoice }]}
        filterFunction={listDisplay}
        onChange={pageChanged}
      />
    </div>
  );
};

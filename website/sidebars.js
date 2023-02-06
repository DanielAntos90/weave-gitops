/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

module.exports = {
  docs: [
    'intro',
    {
      type: 'category',
      label: 'Installation',
      link: {
        type: 'doc',
        id: 'installation/index',
      },
      items: [
        'installation/weave-gitops',
        {
          type: 'category',
          label: 'Weave GitOps Enterprise',
          link: {
            type: 'doc',
            id: 'installation/weave-gitops-enterprise/index',
          },
          items: [
            'installation/weave-gitops-enterprise/airgap',
          ],
        },
        'installation/aws-marketplace',
      ],
    },
    'getting-started',
    {
      type: 'category',
      label: 'Configuration',
      items: [
        'configuration/recommended-rbac-configuration',
        'configuration/securing-access-to-the-dashboard',
        'configuration/service-account-permissions',
        'configuration/user-permissions',
        'configuration/tls',
      ],
    },
    {
      type: 'category',
      label: 'Guides',
      items: [
        'guides/deploying-capa',
        'guides/cert-manager',
        'guides/setting-up-dex',
        'guides/displaying-custom-metadata',
        'guides/using-terraform-templates',
        'guides/delivery',
        'guides/flagger-manual-gating',
      ],
    },
    {
      type: 'category',
      label: 'GitOps Run',
      items: [
        'gitops-run/overview',
        'gitops-run/get-started',
      ],
    },
    {
      type: 'category',
      label: 'Cluster Management',
      items: [
        'cluster-management/intro',
        'cluster-management/getting-started',
        'cluster-management/cluster-api-providers',
        'cluster-management/managing-existing-clusters',
        'cluster-management/provider-identities',
        'cluster-management/deleting-a-cluster',
        'cluster-management/profiles',
        'cluster-management/add-applications',
        'cluster-management/gitrepo-selection',
      ],
    },
    {
      type: 'category',
      label: 'Terraform',
      items: [
        'terraform/overview',
        'terraform/get-started',
        {
          type: 'category',
          label: 'Using Terraform CR',
          items: [
            'terraform/using-terraform-cr/provision-resources-and-auto-approve',
            'terraform/using-terraform-cr/plan-and-manually-apply-terraform-resources',
            'terraform/using-terraform-cr/provision-resources-and-write-output-data',
            'terraform/using-terraform-cr/detect-drifts-only-without-plan-or-apply',
            'terraform/using-terraform-cr/drift-detection-disabled',
            'terraform/using-terraform-cr/set-variables-for-trraform-resources',
            'terraform/using-terraform-cr/custom-backend',
            'terraform/using-terraform-cr/depends-on',
            'terraform/using-terraform-cr/modules',
            'terraform/using-terraform-cr/customize-runner',
          ],
        },
        'terraform/backup-and-restore',
        'terraform/oci-artifact',
        'terraform/aws-eks',
        'terraform/terraform-enterprise',
        'terraform/tfctl',
        'terraform/environment-variables',
      ],
    },
    {
      type: 'category',
      label: 'Pipelines',
      items: [
        'pipelines/intro',
        'pipelines/getting-started',
        'pipelines/authorization',
        'pipelines/promoting-applications',
        'pipelines/pipeline-templates',
        'pipelines/pipeline-templates',
        {
          type: 'category',
          label: 'Reference',
          items: [
            {
              type: 'category',
              label: 'v1alpha1',
              items: [
                'pipelines/spec/v1alpha1/pipeline',
              ],
            },
          ],
        },
      ],
    },
    {
      type: 'category',
      label: 'Workspaces',
      items: [
        'workspaces/intro',
        'workspaces/multi-tenancy',
        'workspaces/view-workspaces',
      ],
    },
    {
      type: 'category',
      label: 'Policy',
      link: {
        type: 'doc',
        id: 'policy/intro',
      },
      items: [
        'policy/getting-started',
        'policy/weave-policy-profile',
        'policy/configuration',
        'policy/policy-set',
        'policy/policy-configuration',
        'policy/releases',
        'policy/commit-time-checks',
      ],
    },
    {
      type: 'category',
      label: 'Templates',
      items: [
        'gitops-templates/templates',
      ],
    },
    'help-and-support',
    'feedback-and-telemetry',
    'releases',
    {
      type: 'category',
      label: 'References',
      items: [
        'references/helm-reference',
        {
          type: 'category',
          label: 'CLI Reference',
          items: [
            {
              type: 'autogenerated',
              dirName: 'references/cli-reference',
            },
          ],
        },
      ],
    },
  ],
};
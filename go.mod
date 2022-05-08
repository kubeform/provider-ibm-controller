module kubeform.dev/provider-ibm-controller

go 1.17

require (
	github.com/IBM-Cloud/terraform-provider-ibm v1.28.1-0.20210729083022-bc68146ed6f3
	github.com/fatih/structs v1.1.0
	github.com/go-logr/logr v0.4.0
	github.com/gobuffalo/flect v0.2.3
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
	github.com/hashicorp/go-uuid v1.0.2
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/terraform-plugin-go v0.3.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/imdario/mergo v0.3.12
	github.com/json-iterator/go v1.1.12
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.3
	go.bytebuilders.dev/audit v0.0.11
	go.bytebuilders.dev/license-verifier v0.9.3
	go.bytebuilders.dev/license-verifier/kubernetes v0.9.2
	gomodules.xyz/logs v0.0.4
	gomodules.xyz/x v0.0.8
	k8s.io/api v0.21.1
	k8s.io/apiextensions-apiserver v0.21.1
	k8s.io/apimachinery v0.21.1
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/klog/v2 v2.8.0
	kmodules.xyz/client-go v0.0.0-20211028120227-48eb36f92a30
	kubeform.dev/apimachinery v0.0.0-20210824104859-ba5604d5a1cc
	kubeform.dev/provider-ibm-api v0.4.1-0.20220308120514-2448e31311de
	kubeform.dev/terraform-backend-sdk v0.0.0-20210922115523-21574335f0db
	sigs.k8s.io/cli-utils v0.25.0
	sigs.k8s.io/controller-runtime v0.9.0
)

require (
	cloud.google.com/go v0.79.0 // indirect
	cloud.google.com/go/storage v1.10.0 // indirect
	github.com/Azure/azure-sdk-for-go v52.5.0+incompatible // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.18 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.13 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.2 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/IBM-Cloud/bluemix-go v0.0.0-20210706092416-dc4055f8e0ea // indirect
	github.com/IBM-Cloud/container-services-go-sdk v0.0.0-20210705152127-41ca00fc9a62 // indirect
	github.com/IBM-Cloud/power-go-client v1.0.55 // indirect
	github.com/IBM/apigateway-go-sdk v0.0.0-20210714141226-a5d5d49caaca // indirect
	github.com/IBM/appconfiguration-go-admin-sdk v0.1.0 // indirect
	github.com/IBM/container-registry-go-sdk v0.0.13 // indirect
	github.com/IBM/go-sdk-core/v3 v3.2.4 // indirect
	github.com/IBM/go-sdk-core/v4 v4.10.0 // indirect
	github.com/IBM/go-sdk-core/v5 v5.5.1 // indirect
	github.com/IBM/ibm-cos-sdk-go v1.7.0 // indirect
	github.com/IBM/ibm-cos-sdk-go-config v1.2.0 // indirect
	github.com/IBM/ibm-hpcs-tke-sdk v0.0.0-20210723145459-a232c3f3ac91 // indirect
	github.com/IBM/keyprotect-go-client v0.7.0 // indirect
	github.com/IBM/networking-go-sdk v0.18.0 // indirect
	github.com/IBM/platform-services-go-sdk v0.18.16 // indirect
	github.com/IBM/push-notifications-go-sdk v0.0.0-20210310100607-5790b96c47f5 // indirect
	github.com/IBM/schematics-go-sdk v0.0.2 // indirect
	github.com/IBM/secrets-manager-go-sdk v0.1.19 // indirect
	github.com/IBM/vpc-go-sdk v0.8.0 // indirect
	github.com/Logicalis/asn1 v0.0.0-20190312173541-d60463189a56 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/PromonLogicalis/asn1 v0.0.0-20190312173541-d60463189a56 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/ScaleFT/sshkeys v0.0.0-20200327173127-6142f742bca5 // indirect
	github.com/Shopify/sarama v1.27.2 // indirect
	github.com/agext/levenshtein v1.2.2 // indirect
	github.com/apache/openwhisk-client-go v0.0.0-20200201143223-a804fb82d105 // indirect
	github.com/apparentlymart/go-cidr v1.1.0 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/apparentlymart/go-versions v1.0.1 // indirect
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/aws/aws-sdk-go v1.40.25 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/bmatcuk/doublestar v1.1.5 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/cloudevents/sdk-go/v2 v2.5.0 // indirect
	github.com/cloudfoundry/jibber_jabber v0.0.0-20151120183258-bcc4c8345a21 // indirect
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dchest/bcrypt_pbkdf v0.0.0-20150205184540-83f37f9c154a // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/emicklei/go-restful v2.9.5+incompatible // indirect
	github.com/evanphx/json-patch v4.11.0+incompatible // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/form3tech-oss/jwt-go v3.2.2+incompatible // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-openapi/analysis v0.19.16 // indirect
	github.com/go-openapi/errors v0.19.9 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/loads v0.20.0 // indirect
	github.com/go-openapi/runtime v0.19.24 // indirect
	github.com/go-openapi/spec v0.20.0 // indirect
	github.com/go-openapi/strfmt v0.20.1 // indirect
	github.com/go-openapi/swag v0.19.12 // indirect
	github.com/go-openapi/validate v0.20.1 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.1+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/gophercloud/gophercloud v0.10.1-0.20200424014253-c3bfe50899e5 // indirect
	github.com/gophercloud/utils v0.0.0-20200423144003-7c72efc7435d // indirect
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	github.com/hashicorp/aws-sdk-go-base v0.7.1 // indirect
	github.com/hashicorp/consul/api v1.9.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-azure-helpers v0.14.0 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-getter v1.5.3 // indirect
	github.com/hashicorp/go-hclog v0.15.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.4.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.6.6 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-safetemp v1.0.0 // indirect
	github.com/hashicorp/go-version v1.3.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl/v2 v2.10.1 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/serf v0.9.5 // indirect
	github.com/hashicorp/terraform-exec v0.14.0 // indirect
	github.com/hashicorp/terraform-json v0.12.0 // indirect
	github.com/hashicorp/terraform-svchost v0.0.0-20200729002733-f050f53b9734 // indirect
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d // indirect
	github.com/hokaccha/go-prettyjson v0.0.0-20170213120834-e6b9231a2b1c // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jcmturner/gofork v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/joyent/triton-go v0.0.0-20180313100802-d8f9c0314926 // indirect
	github.com/jpillora/go-ogle-analytics v0.0.0-20161213085824-14b04e0594ef // indirect
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/klauspost/compress v1.11.2 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.8.0 // indirect
	github.com/likexian/gokit v0.20.15 // indirect
	github.com/lusis/go-artifactory v0.0.0-20160115162124-7e4ce345df82 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/minsikl/netscaler-nitro-go v0.0.0-20170827154432-5b14ce3643e3 // indirect
	github.com/mitchellh/cli v1.1.2 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.0.4 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.4.0 // indirect
	github.com/mitchellh/panicwrap v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mozillazg/go-httpheader v0.3.0 // indirect
	github.com/nats-io/nats.go v1.13.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/nicksnyder/go-i18n v1.10.0 // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/pierrec/lz4 v2.5.2+incompatible // indirect
	github.com/posener/complete v1.2.3 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.18.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/qri-io/jsonpointer v0.1.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/softlayer/softlayer-go v1.0.3 // indirect
	github.com/softlayer/xmlrpc v0.0.0-20200409220501-5f089df7cb7e // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.232 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag v1.0.233 // indirect
	github.com/tencentyun/cos-go-sdk-v5 v0.7.29 // indirect
	github.com/tombuildsstuff/giovanni v0.15.1 // indirect
	github.com/ulikunitz/xz v0.5.8 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v4 v4.3.12 // indirect
	github.com/vmihailenco/tagparser v0.1.1 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/zclconf/go-cty v1.9.1 // indirect
	github.com/zclconf/go-cty-yaml v1.0.2 // indirect
	go.mongodb.org/mongo-driver v1.5.1 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.0 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/oauth2 v0.0.0-20210313182246-cd4f82c27b84 // indirect
	golang.org/x/sys v0.0.0-20210817190340-bfb29a6856f2 // indirect
	golang.org/x/term v0.0.0-20210220032956-6a3ed077a48d // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	golang.org/x/tools v0.1.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gomodules.xyz/clock v0.0.0-20200817085942-06523dba733f // indirect
	gomodules.xyz/flags v0.1.2 // indirect
	gomodules.xyz/homedir v0.1.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.2.0 // indirect
	gomodules.xyz/jsonpath v0.0.1 // indirect
	gomodules.xyz/mergo v0.3.13-0.20210702100041-9d62ff8ece4d // indirect
	gomodules.xyz/pointer v0.1.0 // indirect
	gomodules.xyz/sets v0.2.0 // indirect
	gomodules.xyz/sync v0.1.0 // indirect
	gomodules.xyz/wait v0.0.0-20210218105358-639bd9211545 // indirect
	google.golang.org/api v0.44.0-impersonate-preview // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210319143718-93e7006c17a6 // indirect
	google.golang.org/grpc v1.36.0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/jcmturner/aescts.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/dnsutils.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/gokrb5.v7 v7.5.0 // indirect
	gopkg.in/jcmturner/rpc.v1 v1.1.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
	k8s.io/apiserver v0.21.1 // indirect
	k8s.io/component-base v0.21.1 // indirect
	k8s.io/kube-aggregator v0.21.1 // indirect
	k8s.io/kube-openapi v0.0.0-20210305001622-591a79e4bda7 // indirect
	k8s.io/utils v0.0.0-20210802155522-efc7438f0176 // indirect
	kmodules.xyz/apiversion v0.2.0 // indirect
	kmodules.xyz/custom-resources v0.0.0-20211007080833-72bd9e8cae6e // indirect
	kmodules.xyz/offshoot-api v0.0.0-20210829122105-6f4d481b0c61 // indirect
	kmodules.xyz/resource-metadata v0.6.7 // indirect
	kmodules.xyz/resource-metrics v0.0.5 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.1.2 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace github.com/softlayer/softlayer-go v1.0.3 => github.com/IBM-Cloud/softlayer-go v1.0.3-tf

replace github.com/json-iterator/go => github.com/gomodules/json-iterator v1.1.12-0.20210506053207-2a3ea71074bc

replace bitbucket.org/ww/goautoneg => gomodules.xyz/goautoneg v0.0.0-20120707110453-a547fc61f48d

replace cloud.google.com/go => cloud.google.com/go v0.54.0

replace cloud.google.com/go/bigquery => cloud.google.com/go/bigquery v1.4.0

replace cloud.google.com/go/datastore => cloud.google.com/go/datastore v1.1.0

replace cloud.google.com/go/firestore => cloud.google.com/go/firestore v1.1.0

replace cloud.google.com/go/pubsub => cloud.google.com/go/pubsub v1.2.0

replace cloud.google.com/go/storage => cloud.google.com/go/storage v1.6.0

replace github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d

replace github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible

replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2

replace github.com/golang/protobuf => github.com/golang/protobuf v1.4.3

replace github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.1

replace github.com/imdario/mergo => github.com/imdario/mergo v0.3.5

replace github.com/prometheus-operator/prometheus-operator => github.com/prometheus-operator/prometheus-operator v0.47.0

replace github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring => github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.47.0

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.10.0

replace go.etcd.io/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200910180754-dd1b699fc489

replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20201110150050-8816d57aaa9a

replace google.golang.org/grpc => google.golang.org/grpc v1.32.0

replace helm.sh/helm/v3 => github.com/kubepack/helm/v3 v3.1.0-rc.1.0.20210503022716-7e2d4913a125

replace k8s.io/api => k8s.io/api v0.21.1

replace k8s.io/apimachinery => github.com/kmodules/apimachinery v0.21.2-rc.0.0.20210617231004-332981b97d2d

replace k8s.io/apiserver => github.com/kmodules/apiserver v0.21.2-0.20210617231348-daadbf0c8d5e

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.21.1

replace k8s.io/client-go => k8s.io/client-go v0.21.1

replace k8s.io/component-base => k8s.io/component-base v0.21.1

replace k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20210305001622-591a79e4bda7

replace k8s.io/kubernetes => github.com/kmodules/kubernetes v1.22.0-alpha.0.0.20210617232219-a432af45d932

replace k8s.io/utils => k8s.io/utils v0.0.0-20201110183641-67b214c5f920

replace sigs.k8s.io/application => github.com/kmodules/application v0.8.4-0.20210427030912-90eeee3bc4ad

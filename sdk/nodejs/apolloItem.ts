// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ApolloItem for apolloconfig
 */
export class ApolloItem extends pulumi.CustomResource {
    /**
     * Get an existing ApolloItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApolloItem {
        return new ApolloItem(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'apolloconfig:index:ApolloItem';

    /**
     * Returns true if the given object is an instance of ApolloItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApolloItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApolloItem.__pulumiType;
    }

    /**
     * The appId.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * The clusterName.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * The comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * The dataChangeCreatedBy.
     */
    public readonly dataChangeCreatedBy!: pulumi.Output<string | undefined>;
    /**
     * The dataChangeLastModifiedBy.
     */
    public readonly dataChangeLastModifiedBy!: pulumi.Output<string | undefined>;
    /**
     * The env.
     */
    public readonly env!: pulumi.Output<string>;
    /**
     * The key.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The namespace.
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * The operator for delete item.
     */
    public readonly operator!: pulumi.Output<string | undefined>;
    /**
     * The value.
     */
    public readonly value!: pulumi.Output<string | undefined>;

    /**
     * Create a ApolloItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApolloItemArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.env === undefined) && !opts.urn) {
                throw new Error("Missing required property 'env'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.namespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespace'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dataChangeCreatedBy"] = args ? args.dataChangeCreatedBy : undefined;
            resourceInputs["dataChangeLastModifiedBy"] = args ? args.dataChangeLastModifiedBy : undefined;
            resourceInputs["env"] = args ? args.env : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["operator"] = args ? args.operator : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        } else {
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["clusterName"] = undefined /*out*/;
            resourceInputs["comment"] = undefined /*out*/;
            resourceInputs["dataChangeCreatedBy"] = undefined /*out*/;
            resourceInputs["dataChangeLastModifiedBy"] = undefined /*out*/;
            resourceInputs["env"] = undefined /*out*/;
            resourceInputs["key"] = undefined /*out*/;
            resourceInputs["namespace"] = undefined /*out*/;
            resourceInputs["operator"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["appId", "clusterName", "env", "key", "namespace"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ApolloItem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApolloItem resource.
 */
export interface ApolloItemArgs {
    /**
     * The appId.
     */
    appId: pulumi.Input<string>;
    /**
     * The clusterName.
     */
    clusterName: pulumi.Input<string>;
    /**
     * The comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * The dataChangeCreatedBy.
     */
    dataChangeCreatedBy?: pulumi.Input<string>;
    /**
     * The dataChangeLastModifiedBy.
     */
    dataChangeLastModifiedBy?: pulumi.Input<string>;
    /**
     * The env.
     */
    env: pulumi.Input<string>;
    /**
     * The key.
     */
    key: pulumi.Input<string>;
    /**
     * The namespace.
     */
    namespace: pulumi.Input<string>;
    /**
     * The operator for delete item.
     */
    operator?: pulumi.Input<string>;
    /**
     * The value.
     */
    value?: pulumi.Input<string>;
}

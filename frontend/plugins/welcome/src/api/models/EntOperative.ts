/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntOperativeEdges,
    EntOperativeEdgesFromJSON,
    EntOperativeEdgesFromJSONTyped,
    EntOperativeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntOperative
 */
export interface EntOperative {
    /**
     * 
     * @type {EntOperativeEdges}
     * @memberof EntOperative
     */
    edges?: EntOperativeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntOperative
     */
    id?: number;
    /**
     * OperativeName holds the value of the "operative_Name" field.
     * @type {string}
     * @memberof EntOperative
     */
    operativeName?: string;
}

export function EntOperativeFromJSON(json: any): EntOperative {
    return EntOperativeFromJSONTyped(json, false);
}

export function EntOperativeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntOperative {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntOperativeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'operativeName': !exists(json, 'operative_Name') ? undefined : json['operative_Name'],
    };
}

export function EntOperativeToJSON(value?: EntOperative | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntOperativeEdgesToJSON(value.edges),
        'id': value.id,
        'operative_Name': value.operativeName,
    };
}



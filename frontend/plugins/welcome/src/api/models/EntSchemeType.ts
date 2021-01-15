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
    EntSchemeTypeEdges,
    EntSchemeTypeEdgesFromJSON,
    EntSchemeTypeEdgesFromJSONTyped,
    EntSchemeTypeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSchemeType
 */
export interface EntSchemeType {
    /**
     * SchemeTypeName holds the value of the "SchemeType_Name" field.
     * @type {string}
     * @memberof EntSchemeType
     */
    schemeTypeName?: string;
    /**
     * 
     * @type {EntSchemeTypeEdges}
     * @memberof EntSchemeType
     */
    edges?: EntSchemeTypeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntSchemeType
     */
    id?: number;
}

export function EntSchemeTypeFromJSON(json: any): EntSchemeType {
    return EntSchemeTypeFromJSONTyped(json, false);
}

export function EntSchemeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSchemeType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'schemeTypeName': !exists(json, 'SchemeType_Name') ? undefined : json['SchemeType_Name'],
        'edges': !exists(json, 'edges') ? undefined : EntSchemeTypeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntSchemeTypeToJSON(value?: EntSchemeType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'SchemeType_Name': value.schemeTypeName,
        'edges': EntSchemeTypeEdgesToJSON(value.edges),
        'id': value.id,
    };
}


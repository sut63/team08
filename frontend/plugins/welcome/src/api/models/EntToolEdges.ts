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
    EntOperativerecord,
    EntOperativerecordFromJSON,
    EntOperativerecordFromJSONTyped,
    EntOperativerecordToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntToolEdges
 */
export interface EntToolEdges {
    /**
     * ToolOperativerecord holds the value of the Tool_Operativerecord edge.
     * @type {Array<EntOperativerecord>}
     * @memberof EntToolEdges
     */
    toolOperativerecord?: Array<EntOperativerecord>;
}

export function EntToolEdgesFromJSON(json: any): EntToolEdges {
    return EntToolEdgesFromJSONTyped(json, false);
}

export function EntToolEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntToolEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'toolOperativerecord': !exists(json, 'toolOperativerecord') ? undefined : ((json['toolOperativerecord'] as Array<any>).map(EntOperativerecordFromJSON)),
    };
}

export function EntToolEdgesToJSON(value?: EntToolEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'toolOperativerecord': value.toolOperativerecord === undefined ? undefined : ((value.toolOperativerecord as Array<any>).map(EntOperativerecordToJSON)),
    };
}



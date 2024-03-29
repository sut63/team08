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
 * @interface EntOperativeEdges
 */
export interface EntOperativeEdges {
    /**
     * OperativeOperativerecord holds the value of the Operative_Operativerecord edge.
     * @type {Array<EntOperativerecord>}
     * @memberof EntOperativeEdges
     */
    operativeOperativerecord?: Array<EntOperativerecord>;
}

export function EntOperativeEdgesFromJSON(json: any): EntOperativeEdges {
    return EntOperativeEdgesFromJSONTyped(json, false);
}

export function EntOperativeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntOperativeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'operativeOperativerecord': !exists(json, 'operativeOperativerecord') ? undefined : ((json['operativeOperativerecord'] as Array<any>).map(EntOperativerecordFromJSON)),
    };
}

export function EntOperativeEdgesToJSON(value?: EntOperativeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'operativeOperativerecord': value.operativeOperativerecord === undefined ? undefined : ((value.operativeOperativerecord as Array<any>).map(EntOperativerecordToJSON)),
    };
}



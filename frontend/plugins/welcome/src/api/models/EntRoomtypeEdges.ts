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
    EntRoom,
    EntRoomFromJSON,
    EntRoomFromJSONTyped,
    EntRoomToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRoomtypeEdges
 */
export interface EntRoomtypeEdges {
    /**
     * Rooms holds the value of the rooms edge.
     * @type {Array<EntRoom>}
     * @memberof EntRoomtypeEdges
     */
    rooms?: Array<EntRoom>;
}

export function EntRoomtypeEdgesFromJSON(json: any): EntRoomtypeEdges {
    return EntRoomtypeEdgesFromJSONTyped(json, false);
}

export function EntRoomtypeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRoomtypeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'rooms': !exists(json, 'rooms') ? undefined : ((json['rooms'] as Array<any>).map(EntRoomFromJSON)),
    };
}

export function EntRoomtypeEdgesToJSON(value?: EntRoomtypeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'rooms': value.rooms === undefined ? undefined : ((value.rooms as Array<any>).map(EntRoomToJSON)),
    };
}



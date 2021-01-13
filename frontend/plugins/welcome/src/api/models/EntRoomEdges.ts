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
    EntRent,
    EntRentFromJSON,
    EntRentFromJSONTyped,
    EntRentToJSON,
    EntRoomtype,
    EntRoomtypeFromJSON,
    EntRoomtypeFromJSONTyped,
    EntRoomtypeToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRoomEdges
 */
export interface EntRoomEdges {
    /**
     * 
     * @type {EntRent}
     * @memberof EntRoomEdges
     */
    rents?: EntRent;
    /**
     * 
     * @type {EntRoomtype}
     * @memberof EntRoomEdges
     */
    roomtype?: EntRoomtype;
}

export function EntRoomEdgesFromJSON(json: any): EntRoomEdges {
    return EntRoomEdgesFromJSONTyped(json, false);
}

export function EntRoomEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRoomEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'rents': !exists(json, 'Rents') ? undefined : EntRentFromJSON(json['Rents']),
        'roomtype': !exists(json, 'Roomtype') ? undefined : EntRoomtypeFromJSON(json['Roomtype']),
    };
}

export function EntRoomEdgesToJSON(value?: EntRoomEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'rents': EntRentToJSON(value.rents),
        'roomtype': EntRoomtypeToJSON(value.roomtype),
    };
}



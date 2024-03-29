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

import { exists } from '../runtime';
import {
    EntExaminationroom,
    EntExaminationroomFromJSON,

    EntExaminationroomToJSON,
    EntNurse,
    EntNurseFromJSON,

    EntNurseToJSON,
    EntOperative,
    EntOperativeFromJSON,

    EntOperativeToJSON,
    EntTool,
    EntToolFromJSON,

    EntToolToJSON
} from './';

/**
 * 
 * @export
 * @interface EntOperativerecordEdges
 */
export interface EntOperativerecordEdges {
    /**
     * 
     * @type {EntExaminationroom}
     * @memberof EntOperativerecordEdges
     */
    examinationroom?: EntExaminationroom;
    /**
     * 
     * @type {EntNurse}
     * @memberof EntOperativerecordEdges
     */
    nurse?: EntNurse;
    /**
     * 
     * @type {EntOperative}
     * @memberof EntOperativerecordEdges
     */
    operative?: EntOperative;
    /**
     * 
     * @type {EntTool}
     * @memberof EntOperativerecordEdges
     */
    tool?: EntTool;
}

export function EntOperativerecordEdgesFromJSON(json: any): EntOperativerecordEdges {
    return EntOperativerecordEdgesFromJSONTyped(json, false);
}

export function EntOperativerecordEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntOperativerecordEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'examinationroom': !exists(json, 'Examinationroom') ? undefined : EntExaminationroomFromJSON(json['Examinationroom']),
        'nurse': !exists(json, 'Nurse') ? undefined : EntNurseFromJSON(json['Nurse']),
        'operative': !exists(json, 'Operative') ? undefined : EntOperativeFromJSON(json['Operative']),
        'tool': !exists(json, 'Tool') ? undefined : EntToolFromJSON(json['Tool']),
    };
}

export function EntOperativerecordEdgesToJSON(value?: EntOperativerecordEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'examinationroom': EntExaminationroomToJSON(value.examinationroom),
        'nurse': EntNurseToJSON(value.nurse),
        'operative': EntOperativeToJSON(value.operative),
        'tool': EntToolToJSON(value.tool),
    };
}



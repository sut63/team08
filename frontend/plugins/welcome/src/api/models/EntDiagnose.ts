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
    EntDiagnoseEdges,
    EntDiagnoseEdgesFromJSON,
    EntDiagnoseEdgesFromJSONTyped,
    EntDiagnoseEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDiagnose
 */
export interface EntDiagnose {
    /**
     * DiagnoseID holds the value of the "Diagnose_ID" field.
     * @type {string}
     * @memberof EntDiagnose
     */
    diagnoseID?: string;
    /**
     * DiagnoseNote holds the value of the "Diagnose_Note" field.
     * @type {string}
     * @memberof EntDiagnose
     */
    diagnoseNote?: string;
    /**
     * DiagnoseSymptoms holds the value of the "Diagnose_Symptoms" field.
     * @type {string}
     * @memberof EntDiagnose
     */
    diagnoseSymptoms?: string;
    /**
     * 
     * @type {EntDiagnoseEdges}
     * @memberof EntDiagnose
     */
    edges?: EntDiagnoseEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDiagnose
     */
    id?: number;
}

export function EntDiagnoseFromJSON(json: any): EntDiagnose {
    return EntDiagnoseFromJSONTyped(json, false);
}

export function EntDiagnoseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDiagnose {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'diagnoseID': !exists(json, 'Diagnose_ID') ? undefined : json['Diagnose_ID'],
        'diagnoseNote': !exists(json, 'Diagnose_Note') ? undefined : json['Diagnose_Note'],
        'diagnoseSymptoms': !exists(json, 'Diagnose_Symptoms') ? undefined : json['Diagnose_Symptoms'],
        'edges': !exists(json, 'edges') ? undefined : EntDiagnoseEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntDiagnoseToJSON(value?: EntDiagnose | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Diagnose_ID': value.diagnoseID,
        'Diagnose_Note': value.diagnoseNote,
        'Diagnose_Symptoms': value.diagnoseSymptoms,
        'edges': EntDiagnoseEdgesToJSON(value.edges),
        'id': value.id,
    };
}



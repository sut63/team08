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
/**
 * 
 * @export
 * @interface ControllersDiagnose
 */
export interface ControllersDiagnose {
    /**
     * 
     * @type {number}
     * @memberof ControllersDiagnose
     */
    department?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersDiagnose
     */
    disease?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersDiagnose
     */
    doctor?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersDiagnose
     */
    note?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersDiagnose
     */
    number?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersDiagnose
     */
    patient?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersDiagnose
     */
    symptoms?: string;
}

export function ControllersDiagnoseFromJSON(json: any): ControllersDiagnose {
    return ControllersDiagnoseFromJSONTyped(json, false);
}

export function ControllersDiagnoseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersDiagnose {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'department': !exists(json, 'department') ? undefined : json['department'],
        'disease': !exists(json, 'disease') ? undefined : json['disease'],
        'doctor': !exists(json, 'doctor') ? undefined : json['doctor'],
        'note': !exists(json, 'note') ? undefined : json['note'],
        'number': !exists(json, 'number') ? undefined : json['number'],
        'patient': !exists(json, 'patient') ? undefined : json['patient'],
        'symptoms': !exists(json, 'symptoms') ? undefined : json['symptoms'],
    };
}

export function ControllersDiagnoseToJSON(value?: ControllersDiagnose | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'department': value.department,
        'disease': value.disease,
        'doctor': value.doctor,
        'note': value.note,
        'number': value.number,
        'patient': value.patient,
        'symptoms': value.symptoms,
    };
}



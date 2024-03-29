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
 * @interface ControllersPrescription
 */
export interface ControllersPrescription {
    /**
     * 
     * @type {string}
     * @memberof ControllersPrescription
     */
    added?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersPrescription
     */
    doctor?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPrescription
     */
    drug?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersPrescription
     */
    issue?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersPrescription
     */
    note?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersPrescription
     */
    number?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersPrescription
     */
    nurse?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPrescription
     */
    patient?: number;
}

export function ControllersPrescriptionFromJSON(json: any): ControllersPrescription {
    return ControllersPrescriptionFromJSONTyped(json, false);
}

export function ControllersPrescriptionFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersPrescription {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'added': !exists(json, 'added') ? undefined : json['added'],
        'doctor': !exists(json, 'doctor') ? undefined : json['doctor'],
        'drug': !exists(json, 'drug') ? undefined : json['drug'],
        'issue': !exists(json, 'issue') ? undefined : json['issue'],
        'note': !exists(json, 'note') ? undefined : json['note'],
        'number': !exists(json, 'number') ? undefined : json['number'],
        'nurse': !exists(json, 'nurse') ? undefined : json['nurse'],
        'patient': !exists(json, 'patient') ? undefined : json['patient'],
    };
}

export function ControllersPrescriptionToJSON(value?: ControllersPrescription | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'added': value.added,
        'doctor': value.doctor,
        'drug': value.drug,
        'issue': value.issue,
        'note': value.note,
        'number': value.number,
        'nurse': value.nurse,
        'patient': value.patient,
    };
}


